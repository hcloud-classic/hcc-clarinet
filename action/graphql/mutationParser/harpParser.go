package mutationParser

import (
	"errors"
	"hcc/clarinet/http"
)

func checkSubnetArgsEach(args map[string]interface{}) bool {
	networkIPOk := args["network_ip"].(string) != ""
	netmaskOk := args["netmask"].(string) != ""
	gatewayOk := args["gateway"].(string) != ""
	nextServerOk := args["next_server"].(string) != ""
	nameServerOk := args["name_server"].(string) != ""
	domainNameOk := args["domain_name"].(string) != ""
	serverUUIDOk := args["server_uuid"].(string) != ""
	leaderNodeUUIDOk := args["leader_node_uuid"].(string) != ""
	osOk := args["os"].(string) != ""
	subnetNameOk := args["subnet_name"].(string) != ""

	return networkIPOk || netmaskOk || gatewayOk || nextServerOk || nameServerOk || domainNameOk || serverUUIDOk || leaderNodeUUIDOk || osOk || subnetNameOk
}

func checkSubnetArgsAll(args map[string]interface{}) bool {
	networkIPOk := args["network_ip"].(string) != ""
	netmaskOk := args["netmask"].(string) != ""
	gatewayOk := args["gateway"].(string) != ""
	nextServerOk := args["next_server"].(string) != ""
	nameServerOk := args["name_server"].(string) != ""
	domainNameOk := args["domain_name"].(string) != ""
	serverUUIDOk := args["server_uuid"].(string) != ""
	leaderNodeUUIDOk := args["leader_node_uuid"].(string) != ""
	osOk := args["os"].(string) != ""
	subnetNameOk := args["subnet_name"].(string) != ""

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

	query := "mutation _ { create_subnet(network_ip: \"" + networkIP + "\", netmask: \"" + netmask + "\", gateway: \"" +
		gateway + "\", next_server: \"" + nextServer + "\", name_server: \"" + nameServer + "\", domain_name: \"" +
		domainName + "\", server_uuid: \"" + serverUUID + "\", leader_node_uuid: \"" + leaderNodeUUID + "\", os: \"" +
		os + "\", subnet_name: \"" + subnetName + "\") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", query)
}

func UpdateSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if !checkSubnetArgsEach(args) {
		return nil, errors.New("need some arguments")
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

	arguments := "uuid:\"" + requestedUUID + "\""
	if networkIP != "" {
		arguments += "network_ip:\"" + networkIP + "\","
	}
	if netmask != "" {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gateway != "" {
		arguments += "gateway:\"" + gateway + "\","
	}
	if nextServer != "" {
		arguments += "next_server:\"" + nextServer + "\","
	}
	if nameServer != "" {
		arguments += "next_server:\"" + nameServer + "\","
	}
	if domainName != "" {
		arguments += "domain_name:\"" + domainName + "\","
	}
	if serverUUID != "" {
		arguments += "server_uuid:\"" + serverUUID + "\","
	}
	if leaderNodeUUID != "" {
		arguments += "leader_node_uuid:\"" + leaderNodeUUID + "\","
	}
	if os != "" {
		arguments += "os:\"" + os + "\","
	}
	if subnetName != "" {
		arguments += "subnet_name:\"" + subnetName + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	query := "mutation _ { update_subnet(" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	return http.DoHTTPRequest("harp", query)
}

func DeleteSubnet(args map[string]interface{}) (interface{}, error) {
	requestedUUID, _ := args["uuid"].(string)
	if requestedUUID == "" {
		return nil, errors.New("need a uuid argument")
	}

	query := "mutation _ { delete_subnet(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("harp", query)
}

func CreateDHCPDConf(args map[string]interface{}) (interface{}, error) {
	subnetUUID, _ := args["subnet_uuid"].(string)
	node_uuids, _ := args["node_uuids"].(string)

	if subnetUUID == "" || node_uuids == "" {
		return nil, errors.New("need subnet_uuid and node_uuids arguments")
	}

	query := "mutation _ { create_dhcpd_conf(subnet_uuid: \"" + subnetUUID + "\", node_uuids: \"" + node_uuids + "\") }"

	return http.DoHTTPRequest("harp", query)
}

func checkAdaptiveIPArgsEach(args map[string]interface{}) bool {
	networkAddressOk := args["network_address"].(string) != ""
	netmaskOk := args["netmask"].(string) != ""
	gatewayOk := args["gateway"].(string) != ""
	startIPaddressOk := args["start_ip_address"].(string) != ""
	endIPaddressOk := args["end_ip_address"].(string) != ""

	return networkAddressOk || netmaskOk || gatewayOk || startIPaddressOk || endIPaddressOk
}

func checkAdaptiveIPArgsAll(args map[string]interface{}) bool {
	networkAddressOk := args["network_address"].(string) != ""
	netmaskOk := args["netmask"].(string) != ""
	gatewayOk := args["gateway"].(string) != ""
	startIPaddressOk := args["start_ip_address"].(string) != ""
	endIPaddressOk := args["end_ip_address"].(string) != ""

	return networkAddressOk && netmaskOk && gatewayOk && startIPaddressOk && endIPaddressOk
}

func CreateAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	if !checkAdaptiveIPArgsAll(args) {
		return nil, errors.New("check needed arguments (network_ip, netmask, gateway, next_server, name_server, domain_name, server_uuid, leader_node_uuid, os, adaptiveip_name)")
	}

	networkAddress, _ := args["network_address"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	startIPaddress, _ := args["start_ip_address"].(string)
	endIPaddress, _ := args["end_ip_address"].(string)

	query := "mutation _ { create_adaptiveip(network_address: \"" + networkAddress + "\", netmask: \"" + netmask + "\", gateway: \"" +
		gateway + "\", start_ip_address: \"" + startIPaddress + "\", end_ip_address: \"" + endIPaddress + "\") { uuid network_address netmask gateway start_ip_address end_ip_address } }"

	return http.DoHTTPRequest("harp", query)
}

func UpdateAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a uuid argument")
	}

	if !checkAdaptiveIPArgsEach(args) {
		return nil, errors.New("need some arguments")
	}

	networkIP, _ := args["network_ip"].(string)
	netmask, _ := args["netmask"].(string)
	gateway, _ := args["gateway"].(string)
	startIPaddress, _ := args["start_ip_address"].(string)
	endIPaddress, _ := args["end_ip_address"].(string)

	arguments := "uuid:\"" + requestedUUID + "\""
	if networkIP != "" {
		arguments += "network_ip:\"" + networkIP + "\","
	}
	if netmask != "" {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gateway != "" {
		arguments += "gateway:\"" + gateway + "\","
	}
	if startIPaddress != "" {
		arguments += "start_ip_address:\"" + startIPaddress + "\","
	}
	if endIPaddress != "" {
		arguments += "end_ip_address:\"" + endIPaddress + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	query := "mutation _ { update_adaptiveip(" + arguments + ") { uuid network_address netmask gateway start_ip_address end_ip_address } }"

	return http.DoHTTPRequest("harp", query)
}

func DeleteAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	requestedUUID, _ := args["uuid"].(string)
	if requestedUUID == "" {
		return nil, errors.New("need a uuid argument")
	}

	query := "mutation _ { delete_adaptiveip(uuid:\"" + requestedUUID + "\") { uuid } }"

	return http.DoHTTPRequest("harp", query)
}

func checkAdaptiveIPServerArgsAll(args map[string]interface{}) bool {
	adaptiveIPUUIDOk := args["adaptiveip_uuid"].(string) != ""
	serverUUIDOk := args["server_uuid"].(string) != ""
	publicIPOk := args["public_ip"].(string) != ""

	return adaptiveIPUUIDOk && serverUUIDOk && publicIPOk
}

func CreateAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	if !checkAdaptiveIPServerArgsAll(args) {
		return nil, errors.New("check needed arguments (adaptiveip_uuid, server_uuid, public_ip)")
	}

	adaptiveIPUUID, _ := args["adaptiveip_uuid"].(string)
	serverUUID, _ := args["server_uuid"].(string)
	publicIP, _ := args["public_ip"].(string)

	query := "mutation _ { create_adaptiveip_server(adaptiveip_uuid: \"" + adaptiveIPUUID + "\", server_uuid: \"" + serverUUID + "\", public_ip: \"" +
		publicIP + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway } }"

	return http.DoHTTPRequest("harp", query)
}

func DeleteAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	requestedUUID, requestedUUIDOk := args["server_uuid"].(string)
	if !requestedUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	query := "mutation _ { delete_adaptiveip_server(server_uuid:\"" + requestedUUID + "\") { server_uuid } }"

	return http.DoHTTPRequest("harp", query)
}
