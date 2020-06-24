package queryParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
	"strconv"
)

func Subnet(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	var subnetData data.SubnetData
	query := "query { subnet(uuid: \"" + uuid + "\") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	return http.DoHTTPRequest("harp", true, "SubnetData", subnetData, query)
}

func ListSubnet(args map[string]interface{}) (interface{}, error) {
	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	serverUUID, serverUUIDOk := args["sever_uuid"].(string)
	leaderNodeUUID, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	if networkIPOk {
		arguments += "network_ip:\"" + networkIP + "\","
	}
	if netmaskOk {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gatewayOk {
		arguments += "gateway:\"" + gateway + "\","
	}
	if nextServerOk {
		arguments += "next_server:\"" + nextServer + "\","
	}
	if nameServerOk {
		arguments += "name_server:\"" + nameServer + "\","
	}
	if domainNameOk {
		arguments += "domain_name:\"" + domainName + "\","
	}
	if serverUUIDOk {
		arguments += "sever_uuid:\"" + serverUUID + "\","
	}
	if leaderNodeUUIDOk {
		arguments += "leader_node_uuid:\"" + leaderNodeUUID + "\","
	}
	if osOk {
		arguments += "os:\"" + os + "\","
	}
	if subnetNameOk {
		arguments += "subnet_name:\"" + subnetName + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	var listServerData data.ListServerData
	query := "query { list_subnet(" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	return http.DoHTTPRequest("harp", true, "ListServerData", listServerData, query)
}
