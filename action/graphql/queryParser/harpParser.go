package queryParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
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
	query := `query { ` + cmd + arguments + `{ 
		subnet_list {
			uuid
			network_ip
			netmask
			gateway
			next_server
			name_server
			domain_name
			server_uuid
			leader_node_uuid
			os
			subnet_name
			created_at
		}
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var subnetListData map[string]map[string]model.Subnets
	if e := json.Unmarshal(result, &subnetListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return subnetListData["data"][cmd], nil
}

func ListAdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {

	arguments, err := argumentParser.GetArgumentStr(args)

	cmd := "adaptiveip_available_ip_list"
	query := "query { " + cmd + arguments + " { available_ip_list errors { errtext errcode } } }"

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

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_adaptiveip_server"
	query := `query { ` + cmd + arguments + `{
		adaptiveip_server_list {
			created_at
			private_gateway
			private_ip
			public_ip
			server_uuid
		}
		errors {
			errtext
			errcode
		}
	} }`

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

func AdaptiveIP(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "adaptiveip_setting"
	query := `query { ` + cmd + arguments + `{
		end_ip_address
		ext_ifaceip_address
		gateway_address
		netmask
		start_ip_address
		errors {
			errtext
			errcode
		}
	}}`

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

func ListPortForwarding(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_port_forwarding"
	query := `query { ` + cmd + arguments + `{
		port_forwarding_list {
			server_uuid
			protocol
			external_port
			internal_port
			description
		}
		errors {
			errtext
			errcode
		}
	} }`

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var portForwardingListData map[string]map[string]model.PortForwardingList
	if e := json.Unmarshal(result, &portForwardingListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}

	return portForwardingListData["data"][cmd], nil
}
