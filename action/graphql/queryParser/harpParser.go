package queryParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/lib/errors"
	"hcc/clarinet/model"
)

func ListSubnet(args map[string]string) (interface{}, *errors.HccError) {

	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLArgumentError, "Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_subnet"
	query := "query { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name created_at errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var subnetListData map[string]map[string][]model.Subnet
	if e := json.Unmarshal(result, &subnetListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetListData["data"][cmd], nil
}

func ListAdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {

	cmd := "adaptiveip_available_ip_list"
	query := "query { " + cmd + " { available_ip_list errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipListData map[string]map[string]model.AvailableIPList
	if e := json.Unmarshal(result, &aipListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipListData["data"][cmd], nil
}

func ListAdaptiveIPServer(args map[string]string) (interface{}, *errors.HccError) {
	// server_uuid must checked by cobra
	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLArgumentError, "Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_adaptiveip_server"
	query := "query { " + cmd + arguments + "{ server_uuid public_ip private_ip private_gateway errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipServerListData map[string]map[string]model.AdaptiveIPServers
	if e := json.Unmarshal(result, &aipServerListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipServerListData["data"][cmd], nil
}
