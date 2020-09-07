package mutationParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/lib/errors"
	"hcc/clarinet/model"
)

func CreateSubnet(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args), "server_uuid", "leader_node_uuid"); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_subnet"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name errors } }"
	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]model.Subnet
	if e := json.Unmarshal(result, &subnetData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetData["data"][cmd], nil
}

func UpdateSubnet(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	if argumentParser.CheckArgsMin(args, 2) {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_subnet"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_ip netmask gateway next_server name_server domain_name server_uuid leader_node_uuid os subnet_name errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]model.Subnet
	if e := json.Unmarshal(result, &subnetData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetData["data"][cmd], nil
}

func DeleteSubnet(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_subnet"
	query := "mutation _ { " + cmd + arguments + "{ uuid errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}
	var subnetData map[string]map[string]model.Subnet
	if e := json.Unmarshal(result, &subnetData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetData["data"][cmd], nil
}

func CreateDHCPDConf(args map[string]string) (interface{}, *errors.HccError) {
	// nodeUUID & subnetUUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"subnet_uuid": args["subnet_uuid"],
		"node_uuids":  args["node_uuids"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "create_dhcpd_conf"
	query := "mutation _ { " + cmd + arguments + " { errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var subnetData map[string]map[string]model.DHCPDConfResult
	if e := json.Unmarshal(result, &subnetData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetData["data"][cmd], nil
}

func CreateAdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_address netmask gateway start_ip_address end_ip_address errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipData map[string]map[string]model.AdaptiveIP
	if e := json.Unmarshal(result, &aipData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipData["data"][cmd], nil

}

func UpdateAdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	if argumentParser.CheckArgsMin(args, 2) {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid network_address netmask gateway start_ip_address end_ip_address errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipData map[string]map[string]model.AdaptiveIP
	if e := json.Unmarshal(result, &aipData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipData["data"][cmd], nil
}

func DeleteAdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_adaptiveip"
	query := "mutation _ { " + cmd + arguments + "{ uuid errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipData map[string]map[string]model.AdaptiveIP
	if e := json.Unmarshal(result, &aipData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipData["data"][cmd], nil
}

func CreateAdaptiveIPServer(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_adaptiveip_server"
	query := "mutation _ { " + cmd + arguments + "{ server_uuid public_ip private_ip private_gateway status created_at errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipServerData map[string]map[string]model.AdaptiveIPServer
	if e := json.Unmarshal(result, &aipServerData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipServerData["data"][cmd], nil
}

func DeleteAdaptiveIPServer(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_adaptiveip_server"
	query := "mutation _ { " + cmd + arguments + "{ errors }}"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var aipServerData map[string]map[string]model.AdaptiveIPServer
	if e := json.Unmarshal(result, &aipServerData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return aipServerData["data"][cmd], nil
}
