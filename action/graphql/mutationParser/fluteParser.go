package mutationParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
	"strconv"
)

func checkNodeArgsEach(args map[string]interface{}) bool {
	_, subnetUUIDOk := args["subnet_uuid"].(string)
	_, osOk := args["os"].(string)
	_, serverNameOk := args["server_name"].(string)
	_, serverDescOk := args["server_desc"].(string)
	_, cpuOk := args["cpu"].(int)
	_, memoryOk := args["memory"].(int)
	_, diskSizeOk := args["disk_size"].(int)
	_, statusOk := args["status"].(string)
	_, userUUIDOk := args["user_uuid"].(string)

	return subnetUUIDOk || osOk || serverNameOk || serverDescOk || cpuOk || memoryOk || diskSizeOk || statusOk || userUUIDOk
}

func checkNodeArgsAll(args map[string]interface{}) bool {
	_, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	_, bmcIPOk := args["bmc_ip"].(string)
	_, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	_, statusOk := args["status"].(string)
	_, cpuCoresOk := args["cpu_cores"].(int)
	_, memoryOk := args["memory"].(int)
	_, descriptionOk := args["description"].(string)
	_, activeOk := args["active"].(string)

	return bmcMacAddrOk && bmcIPOk && pxeMacAddrOk && statusOk && cpuCoresOk && memoryOk && descriptionOk && activeOk
}

func checkNodeDetailArgsAll(args map[string]interface{}) bool {
	_, nodeUUIDOk := args["node_uuid"].(string)
	_, cpuModelOk := args["cpu_model"].(string)
	_, cpuProcessorsOk := args["cpu_processors"].(int)
	_, cpuThreadsOk := args["cpu_threads"].(int)

	return nodeUUIDOk && cpuModelOk && cpuProcessorsOk && cpuThreadsOk
}

// Power on Node
func OnOffNode(args map[string]interface{}, state data.PowerState) (interface{}, error) {
	UUID, _ := args["uuid"].(string)
	if UUID == "" {
		return nil, errors.New("need UUID of node")
	}

	query := "mutation _ { "

	switch state {
	case data.On:
		nodeData := data.OnNodeData{}
		query += "on_node(uuid:\"" + UUID + "\") }"
		return http.DoHTTPRequest("flute", true, "OnNodeData", nodeData, query)
	case data.Off:
		nodeData := data.OffNodeData{}
		query += "off_node(uuid:\"" + UUID + "\", force_off: false ) }"
		return http.DoHTTPRequest("flute", true, "OffNodeData", nodeData, query)
	case data.Restart:
		nodeData := data.RestartNodeData{}
		query += "force_restart_node(uuid:\"" + UUID + "\") }"
		return http.DoHTTPRequest("flute", true, "RestartNodeData", nodeData, query)

	}

	return nil, errors.New("Undefined state")
}

func CreateNode(args map[string]interface{}) (interface{}, error) {
	if !checkNodeArgsAll(args) {
		return nil, errors.New("check needed arguments (bmc_mac_addr, bmc_ip, pxe_mac_addr, status, cpu_cores, memory, description, active)")
	}

	bmcMacAddr, _ := args["bmc_mac_addr"].(string)
	bmcIP, _ := args["bmc_ip"].(string)
	pxeMacAddr, _ := args["pxe_mac_addr"].(string)
	status, _ := args["status"].(string)
	cpuCores, _ := args["cpu_cores"].(int)
	memory, _ := args["memory"].(int)
	description, _ := args["description"].(string)
	active, _ := args["active"].(string)

	var createNodeData data.CreateNodeData
	query := "mutation _ { create_node(bmc_mac_addr: \"" + bmcMacAddr + "\", bmc_ip: \"" + bmcIP + "\", pxe_mac_addr: \"" +
		pxeMacAddr + "\", status: \"" + status + "\", cpu_cores: " + strconv.Itoa(cpuCores) + ", memory: " +
		strconv.Itoa(memory) + ", description: \"" + description + "\", active: \"" +
		active + "\") { uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	return http.DoHTTPRequest("flute", true, "CreateNodeData", createNodeData, query)
}

func UpdateNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if checkNodeArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	bmcMacAddr, bmcMacAddrOk := args["bmc_mac_addr"].(string)
	bmcIP, bmcIPOk := args["bmc_ip"].(string)
	pxeMacAddr, pxeMacAddrOk := args["pxe_mac_addr"].(string)
	status, statusOk := args["status"].(string)
	cpuCores, cpuCoresOk := args["cpu_cores"].(int)
	memory, memoryOk := args["memory"].(int)
	description, descriptionOk := args["description"].(string)
	active, activeOk := args["active"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if bmcMacAddrOk {
		arguments += "bmc_mac_addr:\"" + bmcMacAddr + "\","
	}
	if bmcIPOk {
		arguments += "bmc_ip:\"" + bmcIP + "\","
	}
	if pxeMacAddrOk {
		arguments += "pxe_mac_addr:\"" + pxeMacAddr + "\","
	}
	if statusOk {
		arguments += "args:\"" + status + "\","
	}
	if cpuCoresOk {
		arguments += "cpu_cores:" + strconv.Itoa(cpuCores) + ","
	}
	if memoryOk {
		arguments += "memory:" + strconv.Itoa(memory) + "\","
	}
	if descriptionOk {
		arguments += "description:\"" + description + "\","
	}
	if activeOk {
		arguments += "active:\"" + active + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var updateNodeData data.UpdateNodeData
	query := "mutation _ { update_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	return http.DoHTTPRequest("violin", true, "UpdateNodeData", updateNodeData, query)
}

func DeleteNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteNodeData data.DeleteNodeData
	query := "mutation _ { delete_node(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("flute", true, "DeleteNodeData", deleteNodeData, query)
}

func CreateNodeDetail(args map[string]interface{}) (interface{}, error) {
	if !checkNodeDetailArgsAll(args) {
		return nil, errors.New("check needed arguments (node_uuid, cpu_model, cpu_processors, cpu_threads)")
	}

	nodeUUID, _ := args["node_uuid"].(string)
	cpuModel, _ := args["cpu_model"].(string)
	cpuProcessors, _ := args["cpu_processors"].(int)
	cpuThreads, _ := args["cpu_threads"].(int)

	var createNodeDetailData data.CreateNodeDetailData
	query := "mutation _ { create_node_detail(node_uuid: \"" + nodeUUID + "\", cpu_model: \"" + cpuModel +
		"\", cpu_processors: " + strconv.Itoa(cpuProcessors) + ", cpu_threads: " + strconv.Itoa(cpuThreads) +
		") { node_uuid cpu_model cpu_processors cpu_threads } }"

	return http.DoHTTPRequest("flute", true, "CreateNodeDetailData", createNodeDetailData, query)
}

func DeleteNodeDetail(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["node_uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a node_uuid argument")
	}

	var deleteNodeDetailData data.DeleteNodeDetailData
	query := "mutation _ { delete_node_detail(node_uuid:\"" + requestedUUID + "\") { node_uuid } }"

	return http.DoHTTPRequest("flute", true, "DeleteNodeDetailData", deleteNodeDetailData, query)
}
