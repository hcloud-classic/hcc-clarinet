package queryParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
	"hcc/clarinet/model"
)

func ListSubnet(args map[string]string) (interface{}, error) {

	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_subnet"
	query := "query { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var subnetListData map[string]map[string][]model.Subnet
	err = json.Unmarshal(result, &subnetListData)
	if err != nil {
		return nil, err
	}
	return subnetListData["data"][cmd], nil
}

func ListAdaptiveIP(args map[string]string) (interface{}, error) {

	cmd := "adaptiveip_available_ip_list"
	query := "query { " + cmd + " { available_ip_list } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipListData map[string]map[string]map[string][]string
	err = json.Unmarshal(result, &aipListData)
	if err != nil {
		return nil, err
	}
	return aipListData["data"][cmd]["available_ip_list"], nil
}

func ListAdaptiveIPServer(args map[string]string) (interface{}, error) {
	// server_uuid must checked by cobra
	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_adaptiveip_server"
	query := "query { " + cmd + arguments + "{ server_uuid public_ip private_ip private_gateway } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipServerListData map[string]map[string][]model.AdaptiveIPServer
	err = json.Unmarshal(result, &aipServerListData)
	if err != nil {
		return nil, err
	}
	return aipServerListData["data"][cmd], nil
}
