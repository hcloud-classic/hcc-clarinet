package queryParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
	"hcc/clarinet/model"
	"strconv"
)

func Subnet(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "subnet"
	query := "query { " + cmd + " (" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

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

func ListSubnet(args map[string]string) (interface{}, error) {

	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_subnet"
	query := "query { " + cmd + " (" + arguments + ") { uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var subnetListData map[string]map[string][]model.Subnet
	err = json.Unmarshal(result, subnetListData)
	if err != nil {
		return nil, err
	}
	return subnetListData["data"][cmd], nil
}

func AdaptiveIP(args map[string]interface{}) (interface{}, error) {
	uuid, uuidOk := args["uuid"].(string)

	if !uuidOk {
		return nil, errors.New("need a uuid argument")
	}

	query := "query { adaptiveip(uuid: \"" + uuid + "\") { uuid network_addres netmask gateway start_ip_address end_ip_address created_at} }"

	return http.DoHTTPRequest("harp", query)
}

func ListAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	networkAddress, networkAddressOk := args["network_address"].(string)
	netmask, netmaskOk := args["netmask"].(string)
	gateway, gatewayOk := args["gateway"].(string)
	startIPaddress, startIPaddressOk := args["start_ip_address"].(string)
	endIPaddress, endIPaddressOk := args["end_ip_address"].(string)

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	if networkAddressOk {
		arguments += "network_address:\"" + networkAddress + "\","
	}
	if netmaskOk {
		arguments += "netmask:\"" + netmask + "\","
	}
	if gatewayOk {
		arguments += "gateway:\"" + gateway + "\","
	}
	if startIPaddressOk {
		arguments += "start_ip_address:\"" + startIPaddress + "\","
	}
	if endIPaddressOk {
		arguments += "end_ip_address:\"" + endIPaddress + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	query := "query { list_adaptiveip(" + arguments + ") { uuid network_addres netmask gateway start_ip_address end_ip_address created_at} }"

	return http.DoHTTPRequest("harp", query)
}

func AllAdaptiveIP(args map[string]interface{}) (interface{}, error) {
	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_adaptiveip { uuid network_addres netmask gateway start_ip_address end_ip_address created_at} }"
	} else if rowOk && pageOk {
		query = "query { all_adaptiveip(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			") { uuid network_addres netmask gateway start_ip_address end_ip_address created_at} }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	return http.DoHTTPRequest("harp", query)
}

func NumAdaptiveIP() (interface{}, error) {
	query := "query { num_adaptiveip { number } }"

	return http.DoHTTPRequest("harp", query)
}

func AdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	adaptiveIPUUID, adaptiveIPUUIDOk := args["adaptiveip_uuid"].(string)
	serverUUID, serverUUIDOk := args["server_uuid"].(string)

	if !adaptiveIPUUIDOk || !serverUUIDOk {
		return nil, errors.New("need adaptiveip_uuid and server_uuid arguments")
	}

	query := "query { adaptiveip_server(adaptiveip_uuid: \"" + adaptiveIPUUID + "\", server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"

	return http.DoHTTPRequest("harp", query)
}

func ListAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	publicIP, publicIPOk := args["public_ip"].(string)
	privateIP, privateIPOk := args["private_ip"].(string)
	privateGateway, privateGatewayOk := args["private_gateway"].(string)

	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	if !rowOk || !pageOk {
		return nil, errors.New("need row and page arguments")
	}

	arguments := "row:" + strconv.Itoa(row) + ",page:" + strconv.Itoa(page) + ","
	arguments += "server_uuid:\"" + serverUUID + "\","

	if publicIPOk {
		arguments += "public_ip:\"" + publicIP + "\","
	}
	if privateIPOk {
		arguments += "private_ip:\"" + privateIP + "\","
	}
	if privateGatewayOk {
		arguments += "private_gateway:\"" + privateGateway + "\","
	}
	arguments = arguments[0 : len(arguments)-1]

	query := "query { list_adaptiveip_server(" + arguments + ") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"

	return http.DoHTTPRequest("harp", query)
}

func AllAdaptiveIPServer(args map[string]interface{}) (interface{}, error) {
	serverUUID, serverUUIDOk := args["server_uuid"].(string)
	if !serverUUIDOk {
		return nil, errors.New("need a server_uuid argument")
	}

	row, rowOk := args["row"].(int)
	page, pageOk := args["page"].(int)
	var query string

	if !rowOk && !pageOk {
		query = "query { all_adaptiveip_server(server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"
	} else if rowOk && pageOk {
		query = "query { all_adaptiveip_server(row:" + strconv.Itoa(row) + ", page:" + strconv.Itoa(page) +
			", server_uuid: \"" + serverUUID + "\") { adaptiveip_uuid server_uuid public_ip private_ip private_gateway} }"
	} else {
		return nil, errors.New("please insert row and page arguments or leave arguments as empty state")
	}

	return http.DoHTTPRequest("harp", query)
}

func NumAdaptiveIPServer() (interface{}, error) {
	query := "query { num_adaptiveip_server { number } }"

	return http.DoHTTPRequest("harp", query)
}
