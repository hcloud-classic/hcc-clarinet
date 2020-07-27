package mutationParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
	"hcc/clarinet/model"
)

func CreateSubnet(args map[string]string) (interface{}, error) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.New("Check flag value of " + ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_subnet"
	query := "mutation _ { " + cmd + " (" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]model.Subnet
	err = json.Unmarshal(result, &subnetData)
	if err != nil {
		return nil, err
	}
	return subnetData["data"][cmd], nil
}

func UpdateSubnet(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	if argumentParser.CheckArgsMin(args, 2) {
		return nil, errors.New("Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_subnet"
	query := "mutation _ { " + cmd + " (" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]model.Subnet
	err = json.Unmarshal(result, &subnetData)
	if err != nil {
		return nil, err
	}
	return subnetData["data"][cmd], nil
}

func DeleteSubnet(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_subnet"
	query := "mutation _ { " + cmd + " (" + arguments + ") { uuid } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}
	var subnetData map[string]map[string]model.Subnet
	err = json.Unmarshal(result, &subnetData)
	if err != nil {
		return nil, err
	}
	return subnetData["data"][cmd], nil
}

func CreateDHCPDConf(args map[string]string) (interface{}, error) {
	// nodeUUID & subnetUUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"subnet_uuid": args["subnet_uuid"],
		"node_uuids":  args["node_uuids"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "create_dhcpd_conf"
	query := "mutation _ { " + cmd + " (" + arguments + ") }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]string
	err = json.Unmarshal(result, &subnetData)
	if err != nil {
		return nil, err
	}
	return subnetData["data"][cmd], nil
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
