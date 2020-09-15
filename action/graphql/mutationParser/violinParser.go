package mutationParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
	"strconv"
)

func checkServerArgsEach(args map[string]interface{}) bool {
	subnetUUIDOk := args["subnet_uuid"].(string) != ""
	osOk := args["os"].(string) != ""
	serverNameOk := args["server_name"].(string) != ""
	serverDescOk := args["server_desc"].(string) != ""
	cpuOk := args["cpu"].(int) != 0
	memoryOk := args["memory"].(int) != 0
	diskSizeOk := args["disk_size"].(int) != 0
	statusOk := args["status"].(string) != ""
	userUUIDOk := args["user_uuid"].(string) != ""

	return subnetUUIDOk || osOk || serverNameOk || serverDescOk || cpuOk || memoryOk || diskSizeOk || statusOk || userUUIDOk
}

func checkServerArgsAll(args map[string]interface{}) bool {
	subnetUUIDOk := args["subnet_uuid"].(string) != ""
	osOk := args["os"].(string) != ""
	serverNameOk := args["server_name"].(string) != ""
	serverDescOk := args["server_desc"].(string) != ""
	cpuOk := args["cpu"].(int) != 0
	memoryOk := args["memory"].(int) != 0
	diskSizeOk := args["disk_size"].(int) != 0
	userUUIDOk := args["user_uuid"].(string) != ""
	nrNodeOk := args["nr_node"].(int) != 0

	return subnetUUIDOk && osOk && serverNameOk && serverDescOk && cpuOk && memoryOk && diskSizeOk && userUUIDOk && nrNodeOk
}

func CreateServer(args map[string]interface{}) (interface{}, error) {
	if !checkServerArgsAll(args) {
		return nil, errors.New("check needed arguments (subnet_uuid, os, server_name, server_desc, cpu, memory, disk_size, user_uuid, nr_node)")
	}

	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	userUUID, _ := args["user_uuid"].(string)
	nrNode, _ := args["nr_node"].(int)

	var createServerData data.CreateServerData
	query := "mutation _ { create_server(subnet_uuid: \"" + subnetUUID + "\", os: \"" + os + "\", server_name: \"" +
		serverName + "\", server_desc: \"" + serverDesc + "\", cpu: " + strconv.Itoa(cpu) + ", memory: " +
		strconv.Itoa(memory) + ", disk_size: " + strconv.Itoa(diskSize) + ", user_uuid: \"" +
		userUUID + "\", nr_node: " + strconv.Itoa(nrNode) + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size user_uuid } }"

	return http.DoHTTPRequest("violin", true, "CreateServerData", createServerData, query)
}

func UpdateServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if !checkServerArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	status, _ := args["status"].(string)
	userUUID, _ := args["user_uuid"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if subnetUUID != "" {
		arguments += "subnet_uuid:\"" + subnetUUID + "\","
	}
	if os != "" {
		arguments += "os:\"" + os + "\","
	}
	if serverName != "" {
		arguments += "server_name:\"" + serverName + "\","
	}
	if serverDesc != "" {
		arguments += "server_desc:\"" + serverDesc + "\","
	}
	if cpu != 0 {
		arguments += "cpu:" + strconv.Itoa(cpu) + ","
	}
	if memory != 0 {
		arguments += "memory:" + strconv.Itoa(memory) + "\","
	}
	if diskSize != 0 {
		arguments += "disk_size:" + strconv.Itoa(diskSize) + "\","
	}
	if status != "" {
		arguments += "status:\"" + status + "\","
	}
	if userUUID != "" {
		arguments += "user_uuid:\"" + userUUID + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var updateServerData data.UpdateServerData
	query := "mutation _ { update_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	return http.DoHTTPRequest("violin", true, "UpdateServerData", updateServerData, query)
}

func DeleteServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteServerData data.DeleteServerData
	query := "mutation _ { delete_server(uuid:\"" + requestedUUID + "\", status:\"Deleted\") { uuid } }"

	return http.DoHTTPRequest("violin", true, "DeleteServerData", deleteServerData, query)
}

func CreateServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	nodeUUID, nodeUUIDOk := args["node_uuid"].(string)

	if !serverUUIDOk || !nodeUUIDOk {
		return nil, errors.New("need server_uuid and node_uuid arguments")
	}

	var createServerNodeData data.CreateServerNodeData
	query := "mutation _ { create_server_node(server_uuid: \"" + serverUUID + "\", node_uuid: \"" + nodeUUID +
		"\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "CreateServerNodeData", createServerNodeData, query)
}

func DeleteServerNode(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	var deleteServerNodeData data.DeleteServerNodeData
	query := "mutation _ { delete_server_node(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("violin", true, "DeleteServerNodeData", deleteServerNodeData, query)
}
