package queryParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
	"strconv"
)

func Server(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var serverData data.ServerData
	query := "query { server(uuid: \"" + uuid + "\") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ServerData", serverData, query)
}

func ListServer(args map[string]interface{}) (interface{}, error) {

	var arguments = ""

	subnetUUID, _ := args["subnet_uuid"].(string)
	os, _ := args["os"].(string)
	serverName, _ := args["server_name"].(string)
	serverDesc, _ := args["server_desc"].(string)
	cpu, _ := args["cpu"].(int)
	memory, _ := args["memory"].(int)
	diskSize, _ := args["disk_size"].(int)
	status, _ := args["status"].(string)
	userUUID, _ := args["user_uuid"].(string)

	row, _ := args["row"].(int)
	page, _ := args["page"].(int)

	if row != 0 && page != 0 {
		arguments += "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	}
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
	if len(arguments) > 0 {
		arguments = arguments[0 : len(arguments)-1]
	}

	var listServerData data.ListServerData
	query := "query { list_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"
	return http.DoHTTPRequest("violin", true, "ListServerData", listServerData, query)
}

func ServerNode(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var serverNodeData data.ServerNodeData
	query := "query { server_node(uuid: \"" + uuid + "\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ServerNodeData", serverNodeData, query)
}

func ListServerNode(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	var listServerNodeData data.ListServerNodeData
	query := "query { list_server_node(server_uuid: \"" + serverUUID + "\") { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "ListServerNodeData", listServerNodeData, query)
}

func AllServerNode() (interface{}, error) {
	var allServerNodeData data.AllServerNodeData
	query := "query { all_server_node { uuid server_uuid node_uuid created_at } }"

	return http.DoHTTPRequest("violin", true, "AllServerNodeData", allServerNodeData, query)
}

func NumNodesServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	var numNodesServer data.NumNodesServerData
	query := "query { num_nodes_server(server_uuid: \"" + serverUUID + "\") { number } }"

	return http.DoHTTPRequest("violin", true, "NumNodesServerData", numNodesServer, query)
}
