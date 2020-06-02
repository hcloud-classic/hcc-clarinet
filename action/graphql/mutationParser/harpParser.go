package mutationParser

import (
	"errors"
	"hcc/clarinet/data"
	"hcc/clarinet/http"
)

func checkSubnetArgsEach(args map[string]interface{}) bool {
	_, networkIPOk := args["network_ip"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, nextServerOk := args["next_server"].(string)
	_, nameServerOk := args["name_server"].(string)
	_, domainNameOk := args["domain_name"].(string)
	_, serverUUIDOk := args["server_uuid"].(string)
	_, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	_, osOk := args["os"].(string)
	_, subnetNameOk := args["subnet_name"].(string)

	return networkIPOk || netmaskOk || gatewayOk || nextServerOk || nameServerOk || domainNameOk || serverUUIDOk || leaderNodeUUIDOk || osOk || subnetNameOk
}

func checkSubnetArgsAll(args map[string]interface{}) bool {
	_, networkIPOk := args["network_ip"].(string)
	_, netmaskOk := args["netmask"].(string)
	_, gatewayOk := args["gateway"].(string)
	_, nextServerOk := args["next_server"].(string)
	_, nameServerOk := args["name_server"].(string)
	_, domainNameOk := args["domain_name"].(string)
	_, serverUUIDOk := args["server_uuid"].(string)
	_, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	_, osOk := args["os"].(string)
	_, subnetNameOk := args["subnet_name"].(string)

	return networkIPOk && netmaskOk && gatewayOk && nextServerOk && nameServerOk && domainNameOk && serverUUIDOk && leaderNodeUUIDOk && osOk && subnetNameOk
}

func CreateSubnet(args map[string]interface{}) (interface{}, error) {
	if !checkSubnetArgsAll(args) {
		return nil, errors.New("check needed arguments (network_ip, netmask, gateway, next_server, name_server, domain_name, server_uuid, leader_node_uuid, os, subnet_name)")
	}

	networkIP, _ := args["network_ip"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	nextServer, _ := args["next_server"].(string)
	nameServer, _ := args["name_server"].(string)
	domainName, _ := args["domain_name"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	leaderNodeUUID, _ := args["leader_node_uuid"].(string)
	os, _ := args["os"].(string)
	subnetName, _ := args["subnet_name"].(string)

	var createSubnetData data.CreateSubnetData
	query := "mutation _ { create_subnet(network_ip: \"" + networkIP + "\", netmask: \"" + netmask + "\", gateway: \"" +
		gateway + "\", next_server: \"" + nextServer + "\", name_server: \"" + nameServer + "\", domain_name: \"" +
		domainName + "\", server_uuid: \"" + serverUUID + "\", leader_node_uuid: \"" + leaderNodeUUID + "\", os: \"" +
		os + "\", subnet_name: \"" + subnetName + "\") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", true, "CreateSubnetData", createSubnetData, query)
}

func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if checkSubnetArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	networkIP, networkIPOk := args["network_ip"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	nextServer, nextServerOk := args["next_server"].(string)
	nameServer, nameServerOk := args["name_server"].(string)
	domainName, domainNameOk := args["domain_name"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	leaderNodeUUID, leaderNodeUUIDOk := args["leader_node_uuid"].(string)
	os, osOk := args["os"].(string)
	subnetName, subnetNameOk := args["subnet_name"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
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
		arguments += "next_server:\"" + nameServer + "\","
	}
	if domainNameOk {
		arguments += "domain_name:\"" + domainName + "\","
	}
	if serverUUIDOk {
		arguments += "server_uuid:\"" + serverUUID + "\","
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

	var updateSubnetData data.UpdateSubnetData
	query := "mutation _ { update_subnet(" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", true, "UpdateSubnetData", updateSubnetData, query)
}
