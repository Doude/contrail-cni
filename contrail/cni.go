// vim: tabstop=4 expandtab shiftwidth=4 softtabstop=4
//
// Copyright (c) 2017 Juniper Networks, Inc. All rights reserved.
//
package contrailCni

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	cniIntf "github.com/michaelhenkel/contrail-cni/common"
	log "github.com/michaelhenkel/contrail-cni/logging"
)

// CNIVersion is the version from the network configuration
var CNIVersion string

/* Example configuration file
{
    "cniVersion": "0.3.1",
    "contrail" : {
        "vrouter-ip"    : "127.0.0.1",
        "vrouter-port"  : 9092,
        "config-dir"    : "/var/lib/contrail/ports/vm",
        "poll-timeout"  : 15,
        "poll-retries"  : 5,
        "log-dir"       : "/var/log/contrail/cni",
        "log-level"     : "2",
        "mode"          : "k8s/mesos",
        "vif-type"      : "veth/macvlan",
        "parent-interface" : "eth0",
        "mtu"           : 1200,
        "meta-plugin"   : false
    },

    "name": "contrail",
    "type": "contrail"
}
*/

const LOG_DIR = "/var/log/contrail/cni"
const LOG_LEVEL = "2"

// Container orchestrator modes
const CNI_MODE_K8S = "k8s"
const CNI_MODE_MESOS = "mesos"

// Type of virtual interface to be created for container
const VIF_TYPE_VETH = "veth"
const VIF_TYPE_MACVLAN = "macvlan"
const VIF_TYPE_ETH = "eth"

// In case of macvlan, the container interfaces will run as sub-interface
// to interface on host network-namespace. Name of the interface inside
// host network-namespace is defined below
const CONTRAIL_PARENT_INTERFACE = "eth0"

// Definition of Logging arguments in form of json in STDIN
type ContrailCni struct {
	cniArgs       *skel.CmdArgs
	Mode          string `json:"mode"`
	VifType       string `json:"vif-type"`
	VifParent     string `json:"parent-interface"`
	LogDir        string `json:"log-dir"`
	LogFile       string `json:"log-file"`
	LogLevel      string `json:"log-level"`
	Mtu           int    `json:"mtu"`
	MetaPlugin    bool   `json:"meta-plugin"`
	NetworkName   string `json:"network-name"`
	ContainerUuid string
	ContainerName string
	ContainerVn   string
	MesosIP       string `json:"mesos-ip"`
	MesosPort     string `json:"mesos-port"`
	VRouter       VRouter
}

type cniJson struct {
	ContrailCni ContrailCni `json:"contrail"`
	CniVersion  string      `json:"cniVersion"`
}

// Apply logging configuration. We use log packet for logging.
// log supports log-dir and log-level as arguments only.
func (cni *ContrailCni) loggingInit() error {
	log.Init(cni.LogFile, 10, 5)
	//flag.Parse()
	//flag.Lookup("log_dir").Value.Set(cni.LogDir)
	//flag.Lookup("v").Value.Set(cni.LogLevel)
	return nil
}

func (cni *ContrailCni) Log() {
	log.Infof("ContainerID : %s\n", cni.cniArgs.ContainerID)
	log.Infof("NetNS : %s\n", cni.cniArgs.Netns)
	log.Infof("Container Ifname : %s\n", cni.cniArgs.IfName)
	log.Infof("Args : %s\n", cni.cniArgs.Args)
	log.Infof("CNI VERSION : %s\n", CNIVersion)
	log.Infof("MTU : %d\n", cni.Mtu)
	log.Infof("Meta Plugin: %t\n", cni.MetaPlugin)
	log.Infof("Network Name: %t\n", cni.NetworkName)
	log.Infof("Config File : %s\n", cni.cniArgs.StdinData)
	log.Infof("%+v\n", cni)
	cni.VRouter.Log()
}

func Init(args *skel.CmdArgs) (*ContrailCni, error) {
	vrouter, _ := VRouterInit(args.StdinData)

	cni := ContrailCni{cniArgs: args, Mode: CNI_MODE_K8S,
		VifType: VIF_TYPE_VETH, VifParent: CONTRAIL_PARENT_INTERFACE,
		LogDir: LOG_DIR, LogLevel: LOG_LEVEL, Mtu: cniIntf.CNI_MTU,
		MetaPlugin: false, VRouter: *vrouter}
	json_args := cniJson{ContrailCni: cni}

	if err := json.Unmarshal(args.StdinData, &json_args); err != nil {
		log.Errorf("Error decoding stdin\n %s \n. Error %+v",
			string(args.StdinData), err)
		return nil, err
	}

	// If CNI version is blank, set to "0.2.0"
	CNIVersion = json_args.CniVersion
	if CNIVersion == "" {
		CNIVersion = "0.2.0"
	}

	json_args.ContrailCni.loggingInit()
	return &json_args.ContrailCni, nil
}

func (cni *ContrailCni) Update(containerName, containerUuid,
	containerVn string) {
	log.Infof("got cni update request with uuid %s name %s vn %s", containerUuid, containerName, containerVn)
	cni.ContainerUuid = containerUuid
	cni.ContainerName = containerName
	cni.ContainerVn = containerVn
	log.Infof("processed cni update request with uuid %s name %s vn %s", cni.ContainerUuid, cni.ContainerName, containerVn)

}

/*
 *  makeInterface- Method to intialize interface object of type VETh or MACVLAN
 */
func (cni *ContrailCni) makeInterface(
	vlanId int, containerIntfName string) cniIntf.CniIntfMethods {
	if cni.VifType == VIF_TYPE_MACVLAN {
		return cniIntf.CniIntfMethods(cniIntf.InitMacVlan(cni.VifParent,
			containerIntfName, cni.cniArgs.ContainerID, cni.ContainerUuid,
			cni.cniArgs.Netns, cni.Mtu, vlanId))
	}

	return cniIntf.CniIntfMethods(cniIntf.InitVEth(containerIntfName,
		cni.cniArgs.ContainerID, cni.ContainerUuid, cni.cniArgs.Netns, cni.Mtu))
}

/*
 * buildContainerIntfName - Method to construct interface name for container.
 * - When Contrail CNI Plugin is used as a meta plugin, interface names are
 *   are generated by appending the provided 'index' to "eth".
 * - If the case when CNI plugin is invoked by a delegating plugin,
 *   the provided IfName arg is returned.
 */
func (cni *ContrailCni) buildContainerIntfName(
	index int, isMetaPlugin bool) string {
	var intfName string
	if isMetaPlugin == true {
		intfName = cni.cniArgs.IfName
	} else {
		intfName = VIF_TYPE_ETH + strconv.Itoa(index)
	}
	log.Infof("Built container interface name - %s", intfName)
	return intfName
}

/****************************************************************************
 * Add message handlers
 ****************************************************************************/

/* createInterfaceAndUpdateVrouter -
 *   Method to create interface in a container and notify VRouter about it
 */
func (cni *ContrailCni) createInterfaceAndUpdateVrouter(
	containerIntfName string, result Result) error {
	intf := cni.makeInterface(result.VlanId, containerIntfName)
	intf.Log()

	err := intf.Create()
	if err != nil {
		log.Errorf("Error creating interface object. Error %+v", err)
		return err
	}

	// Inform vrouter about interface-add
	// The interface inside container must be created by this time.
	updateAgent := true
	if cni.VifType == VIF_TYPE_MACVLAN {
		updateAgent = false
	}

	err = cni.VRouter.Add(cni.ContainerName, cni.ContainerUuid,
		result.VnId, cni.cniArgs.ContainerID, cni.cniArgs.Netns,
		containerIntfName, intf.GetHostIfName(), result.VmiUuid, updateAgent)
	if err != nil {
		log.Errorf("Error in Add to VRouter. Error %+v", err)
		return err
	}

	return nil
}

/* configureContainerInterface -
 *   Method to configure the interface in a container
 */
func (cni *ContrailCni) configureContainerInterface(
	containerIntfName string, vRouterResult Result) (*current.Result, error) {
	containerInterface := cni.makeInterface(
		vRouterResult.VlanId, containerIntfName)
	typesResult := MakeCniResult(containerIntfName, &vRouterResult)

	// Configure the interface based on config received above
	err := containerInterface.Configure(vRouterResult.Mac, typesResult)
	if err != nil {
		log.Errorf("Error configuring container interface. Error %+v", err)
		return nil, err
	}

	return typesResult, nil
}

/*
 *  CmdAdd - Method to
 *  - ADD handler for a container
 *  - Pre-fetch interface configuration from VRouter.
 *  - Gets MAC address for the interface
 *  - In case of sub-interface, gets VLAN-Tag for the interface
 *  - Create interface based on the "mode"
 *  - Invoke Add handler from VRouter module
 *  - Update interface with configuration got from VRouter
 *  - Configures IP address
 *  - Configures routes
 *  - Bring-up the interface
 *  - Return result in form of types.Result
 */
func (cni *ContrailCni) CmdAdd() (*current.Result, string, error) {

	// ContainerIntfNames - map of vmi uuid to interface names
	// Key : vmi uuid
	// Value : interface name
	res := &current.Result{}
	containerIntfNames := make(map[string]string)
	// vrouterResultMap - map of vmi uuid to Vrouter Result
	// Key : vmi uuid
	// Value : VRouter result
	vrouterResultMap := make(map[string]Result)

	var finalTypesResult *current.Result

	if cni.Mode == CNI_MODE_MESOS {
		op := "POST"
		url := "http://" + cni.MesosIP + ":" + cni.MesosPort + "/" + "add_cni_info"
		values := map[string]string{"cid": cni.cniArgs.ContainerID, "cmd": "ADD", "args": string(cni.cniArgs.StdinData)}
		jsonValue, _ := json.Marshal(values)
		log.Infof("IN CNI MESOS mode - Updating Mesos Manager: URL -  %s\n", url)
		req, err := http.NewRequest(op, url, bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Errorf("Error creating http Request. Op %s Url %s Msg %s."+
				"Error : %+v", op, url, err)
			return res, "", err
		}

		req.Header.Set("Content-Type", "application/json")
		httpClient := new(http.Client)
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("Failed HTTP operation :  %+v. Error : %+v", req, err)
			return res, "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			msg := fmt.Sprintf("Failed HTTP POST operation. Return code %d",
				int(resp.StatusCode))
			log.Errorf(msg)
			return res, "", fmt.Errorf(msg)
		}
		log.Infof("IN CNI MESOS mode - Post to mesos manager success!\n")
	}

	// Pre-fetch initial configuration for the interfaces from vrouter
	// This will give MAC address for the interface and in case of
	// VMI sub-interface, we will also get the vlan-tag
	results, err := cni.VRouter.Poll(cni.ContainerUuid, cni.ContainerVn)
	if err != nil {
		log.Errorf("Error polling for configuration of %s and %s",
			cni.ContainerUuid, cni.ContainerVn)
		return res, "", err
	}

	// For each interface in the result create an interface in the container,
	// persist the configuration and notify the Vrouter agent
	for _, result := range *results {
		// Index annotation is expected in the following format -
		//         <index>/<total num of interfaces>
		indices := strings.Split(result.Annotations.Index, "/")
		index, err := strconv.Atoi(indices[0])
		if err != nil {
			log.Errorf("Could not retrieve index from result - %+v", result)
			return res, "", err
		}
		var containerIntfName string
		if result.Annotations.Interface != "" {
			containerIntfName = result.Annotations.Interface
		} else {
			containerIntfName = cni.buildContainerIntfName(index, cni.MetaPlugin)
		}

		if cni.MetaPlugin {
			// When invoked by a delegating plugin, work on the vrouter result that
			// matches the given network name. In the scenario when multiple results
			// match the nw name, check if a config file exists with the VMI UUID
			// from the result. If the file exists look for other result.
			if result.Annotations.Network == cni.NetworkName {
				// Check if this VMI uuid is already used
				fname := cni.VRouter.makeFileName(result.VmiUuid)
				if checkFileOrDirExists(fname) {
					continue
				}
			} else {
				// TODO: Throw error when nw name not in the list of interfaces
				continue
			}
		}

		log.Infof("Creating interface - %s for result - %v",
			containerIntfName, result)
		cni.createInterfaceAndUpdateVrouter(containerIntfName, result)
		containerIntfNames[result.VmiUuid] = containerIntfName

		if cni.MetaPlugin {
			// When invoked by a delegating plugin, Work only on the given interface
			// name and ignore the rest of the interfaces
			break
		}
	}

	vRouterResults, poll_err := cni.VRouter.PollUrl("/vm")
	if poll_err != nil {
		log.Errorf("Error in polling VRouter ")
		return res, "", poll_err
	}

	for _, vRouterResult := range *vRouterResults {
		vrouterResultMap[vRouterResult.VmiUuid] = vRouterResult
	}

	log.Infof("About to configure %d interfaces for container",
		len(containerIntfNames))
	for vmiUuid := range containerIntfNames {
		containerIntfName := containerIntfNames[vmiUuid]
		vRouterResult, ok := vrouterResultMap[vmiUuid]
		if ok == false {
			msg := fmt.Sprintf("VMI UUID %s does not exist in the Vrouter Result",
				vmiUuid)
			log.Errorf(msg)
			return res, "", fmt.Errorf(msg)
		}

		log.Infof("Working on VrouterResult - %+v  and Interface name - %s",
			vRouterResult, containerIntfName)
		typesResult, err := cni.configureContainerInterface(
			containerIntfName, vRouterResult)
		if err != nil {
			return res, "", err
		}

		if cni.MetaPlugin {
			finalTypesResult = typesResult
		} else {
			if containerIntfName == "eth0" {
				finalTypesResult = typesResult
			}
		}
	}

	types.PrintResult(finalTypesResult, CNIVersion)
	return finalTypesResult, CNIVersion, nil
}

/****************************************************************************
 * Delete message handlers
 ****************************************************************************/
func (cni *ContrailCni) CmdDel() error {
	if cni.Mode == CNI_MODE_MESOS {
		op := "POST"
		url := "http://" + cni.MesosIP + ":" + cni.MesosPort + "/" + "del_cni_info"

		values := map[string]string{"cid": cni.cniArgs.ContainerID, "cmd": "DEL",
			"args": string(cni.cniArgs.StdinData)}
		jsonValue, _ := json.Marshal(values)
		log.Infof("IN CNI MESOS mode - Updating Mesos Manager: URL -  %s\n", url)
		req, err := http.NewRequest(op, url, bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Errorf("Error creating http Request. Op %s Url %s Msg %s."+
				"Error : %+v", op, url, err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		httpClient := new(http.Client)
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("Failed HTTP operation :  %+v. Error : %+v", req, err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			msg := fmt.Sprintf("Failed HTTP POST operation. Return code %d",
				int(resp.StatusCode))
			log.Errorf(msg)
			return fmt.Errorf(msg)
		}

		log.Infof("IN CNI MESOS mode - Post to mesos manager success!\n")
	}

	containerIntfNames, vmiUuids, err := cni.VRouter.CanDelete(
		cni.cniArgs.ContainerID, cni.ContainerUuid, cni.ContainerVn)
	if err != nil {
		log.Errorf("Failed in CanDelete. Error %s", err)
		return nil
	}

	if len(containerIntfNames) > 0 {
		for _, containerIntfName := range containerIntfNames {
			intf := cni.makeInterface(0, containerIntfName)
			intf.Log()

			// Build CNI response from response
			err = intf.Delete()
			if err != nil {
				log.Errorf("Error deleting interface")
			} else {
				log.Infof("Deleted interface %s inside container",
					containerIntfName)
			}
		}

		// Inform vrouter about interface-delete.
		updateAgent := true
		if cni.VifType == VIF_TYPE_MACVLAN {
			updateAgent = false
		}

		err = cni.VRouter.Del(cni.cniArgs.ContainerID, cni.ContainerUuid,
			cni.ContainerVn, updateAgent, vmiUuids)
		if err != nil {
			log.Errorf("Error deleting interface from agent")
		}
	}

	// Nothing to do
	return nil
}

func (cni *ContrailCni) CmdCheck() error {
	return nil
}
