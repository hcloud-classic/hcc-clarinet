package mutationParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
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

func OnNode(args map[string]interface{}) (interface{}, error) {
	mac, macOk := args["mac"].(string)
	if !macOk {
		return nil, errors.New("need a mac argument")
	}

	var onNodeData data.OnNodeData
	query := "mutation _ { on_node(mac:\"" + mac + "\") }"

	return http.DoHTTPRequest("flute", true, "OnNodeData", onNodeData, query)
}
