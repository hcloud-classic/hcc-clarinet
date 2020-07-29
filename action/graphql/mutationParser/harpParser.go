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
	query := "mutation _ { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

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
	query := "mutation _ { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name } }"

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
	query := "mutation _ { " + cmd + arguments + "{ uuid } }"

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
	query := "mutation _ { " + cmd + arguments + "}"

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

func CreateAdaptiveIP(args map[string]string) (interface{}, error) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.New("Check flag value of " + ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_address netmask gateway start_ip_address end_ip_address } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipData map[string]map[string]model.AdaptiveIP
	err = json.Unmarshal(result, &aipData)
	if err != nil {
		return nil, err
	}
	return aipData["data"][cmd], nil

}

func UpdateAdaptiveIP(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	if argumentParser.CheckArgsMin(args, 2) {
		return nil, errors.New("Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_address netmask gateway start_ip_address end_ip_address } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipListData map[string]map[string][]model.AdaptiveIP
	err = json.Unmarshal(result, &aipListData)
	if err != nil {
		return nil, err
	}
	return aipListData["data"][cmd], nil
}

func DeleteAdaptiveIP(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid } }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipData map[string]map[string]model.AdaptiveIP
	err = json.Unmarshal(result, &aipData)
	if err != nil {
		return nil, err
	}
	return aipData["data"][cmd], nil
}

func CreateAdaptiveIPServer(args map[string]string) (interface{}, error) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.New("Check flag value of " + ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_adaptiveip_server"
	query := "mutation _ { " + cmd + arguments + "{ server_uuid public_ip private_ip private_gateway status created_at} }"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipServerData map[string]map[string]model.AdaptiveIPServer
	err = json.Unmarshal(result, &aipServerData)
	if err != nil {
		return nil, err
	}
	return aipServerData["data"][cmd], nil
}

func DeleteAdaptiveIPServer(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_adaptiveip_server"
	query := "mutation _ { " + cmd + arguments + "}"

	result, err := http.DoHTTPRequest("harp", query)
	if err != nil {
		return nil, err
	}

	var aipServerData map[string]map[string]model.AdaptiveIPServer
	err = json.Unmarshal(result, &aipServerData)
	if err != nil {
		return nil, err
	}
	return aipServerData["data"][cmd], nil
}
