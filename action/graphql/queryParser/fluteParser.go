package queryParser

import (
	"encoding/json"

	argumentParser "hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

func Node(args map[string]string) (interface{}, *errors.HccError) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "node"
	query := `query { ` + cmd + arguments + `{
		uuid
		bmc_mac_addr
		bmc_ip
		pxe_mac_addr
		status
		cpu_cores
		memory
		description
		created_at
		active
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string]model.Node
	if e := json.Unmarshal(result, &nodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return nodeData["data"][cmd], nil
}

func ListNode(args map[string]string) (interface{}, *errors.HccError) {

	if (args["row"] != "0") != (args["page"] != "0") {

		return nil, errors.NewHccError(errors.ClarinetGraphQLArgumentError, "Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_node"
	query := `query { ` + cmd + arguments + `{
		node_list {
			uuid
			server_uuid
			bmc_mac_addr
			bmc_ip
			pxe_mac_addr
			status
			cpu_cores
			memory
			created_at
			active
		}
		errors {
			errcode
			errtext
		}
	} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string]model.Nodes
	if e := json.Unmarshal(result, &nodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, e.Error())
	}
	return nodeData["data"][cmd], nil
}

// Not Used
func NumNode() (interface{}, *errors.HccError) {
	cmd := "num_node"
	query := "query { " + cmd + " { number errors } }"

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var nodeNum map[string]map[string]model.NodeNum
	if e := json.Unmarshal(result, &nodeNum); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return nodeNum["data"][cmd], nil
}

func NodeDetail(args map[string]string) (interface{}, *errors.HccError) {
	// node_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "detail_node"
	query := `query { ` + cmd + arguments + `{
		node_detail_data
		nic_detail_data
		node_uuid
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var nodeDetail map[string]map[string]model.NodeDetail
	if e := json.Unmarshal(result, &nodeDetail); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}

	var nodeDetailData model.NodeDetailData
	if e := json.Unmarshal([]byte(nodeDetail["data"][cmd].NodeDetail), &nodeDetailData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}

	if e := json.Unmarshal([]byte(nodeDetail["data"][cmd].NicDetail), &nodeDetailData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	// check graphql err
	nodeDetailData.Errors = nodeDetail["data"][cmd].Errors
	nodeDetailData.NodeUUID = nodeDetail["data"][cmd].NodeUUID

	return nodeDetailData, nil
}
