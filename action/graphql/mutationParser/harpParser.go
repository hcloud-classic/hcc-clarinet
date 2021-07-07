package mutationParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

func CreateSubnet(args map[string]string) (interface{}, *errors.HccError) {

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_subnet"
	query := `mutation _ { ` + cmd + arguments + `{
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
		errors {
			errcode
			errtext
		}
	} }`
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
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_subnet"
	query := `mutation _ { ` + cmd + arguments + `{
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
		errors {
			errcode
			errtext
		}
	} }`

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
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_subnet"
	query := `mutation _ { ` + cmd + arguments + `{
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
		errors {
			errcode
			errtext
		}
	} }`

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
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_dhcpd_conf"
	query := `mutation _ { ` + cmd + arguments + `{
		result
		errors {
			errcode
			errtext
		}
	} }`

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

func CreateAdaptiveIPSetting(args map[string]string) (interface{}, *errors.HccError) {
	// All argumets must checked by cobra

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_adaptiveip_setting"
	query := `mutation _ { ` + cmd + arguments + `{
		ext_ifaceip_address
		netmask
		gateway_address
		start_ip_address
		end_ip_address
		errors {
			errcode
			errtext
		}
	} }`

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
	query := `mutation _ { ` + cmd + arguments + `{
		server_uuid
		public_ip
		private_ip
		private_gateway
		created_at
		errors {
			errcode
			errtext
		}
	} }`

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
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_adaptiveip_server"
	query := `mutation _ { ` + cmd + arguments + `{
		server_uuid
		public_ip
		private_ip
		private_gateway
		created_at
		errors {
			errcode
			errtext
		}
	} }`

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

func CreatePortForwarding(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_port_forwarding"
	query := `mutation _ { ` + cmd + arguments + `{
		server_uuid
		protocol
		external_port
		internal_port
		description
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var portForwardingData map[string]map[string]model.PortForwarding
	if e := json.Unmarshal(result, &portForwardingData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}

	return portForwardingData["data"][cmd], nil
}

func DeletePortForwarding(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_port_forwarding"
	query := `mutation _ { ` + cmd + arguments + `{
		server_uuid
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var portForwardingData map[string]map[string]model.PortForwarding
	if e := json.Unmarshal(result, &portForwardingData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return portForwardingData["data"][cmd], nil
}
