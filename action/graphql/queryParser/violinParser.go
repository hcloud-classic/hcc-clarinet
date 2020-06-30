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
	subnetUUID, subnetUUIDOk := args["subnet_uuid"].(string)
	os, osOk := args["os"].(string)
	serverName, serverNameOk := args["server_name"].(string)
	serverDesc, serverDescOk := args["server_desc"].(string)
	cpu, cpuOk := args["cpu"].(int)
	memory, memoryOk := args["memory"].(int)
	diskSize, diskSizeOk := args["disk_size"].(int)
	status, statusOk := args["status"].(string)
	userUUID, userUUIDOk := args["user_uuid"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	if subnetUUIDOk {
		arguments += "subnet_uuid:\"" + subnetUUID + "\","
	}
	if osOk {
		arguments += "os:\"" + os + "\","
	}
	if serverNameOk {
		arguments += "server_name:\"" + serverName + "\","
	}
	if serverDescOk {
		arguments += "server_desc:\"" + serverDesc + "\","
	}
	if cpuOk {
		arguments += "cpu:" + strconv.Itoa(cpu) + ","
	}
	if memoryOk {
		arguments += "memory:" + strconv.Itoa(memory) + "\","
	}
	if diskSizeOk {
		arguments += "disk_size:" + strconv.Itoa(diskSize) + "\","
	}
	if statusOk {
		arguments += "status:\"" + status + "\","
	}
	if userUUIDOk {
		arguments += "user_uuid:\"" + userUUID + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var listServerData data.ListServerData
	query := "query { list_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	return http.DoHTTPRequest("violin", true, "ListServerData", listServerData, query)
}
