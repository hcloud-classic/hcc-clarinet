package mutationParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/lib/errors"
	"hcc/clarinet/model"
)

// Power on Node
func OnOffNode(args map[string]string, state model.PowerState) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	var cmd string
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	query := "mutation _ { "

	switch state {
	case model.On:
		cmd = "on_node"
		query += cmd + arguments + "}"
	case model.Off:
		cmd = "off_node"
		query += cmd + arguments + "}"
	case model.Restart:
		cmd = "force_restart_node"
		query += cmd + arguments + "}"
	default:
		return nil, errors.NewHccError(errors.ClarinetGraphQLRequestError, "Undefined Power state")
	}

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string]string
	if e := json.Unmarshal(result, &nodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return nodeData["data"][cmd], nil
}

func CreateNode(args map[string]string) (interface{}, *errors.HccError) {
	// bmc_ip & description must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	if e := json.Unmarshal(result, &node); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return node["data"][cmd], nil
}

func UpdateNode(args map[string]string) (interface{}, *errors.HccError) {
	// bmc_ip must checked by cobra
	if argumentParser.CheckArgsMin(args, 2, "bmc_ip") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Need at least 1 more flag except bmc_ip")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid bmc_mac_addr errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	if e := json.Unmarshal(result, &node); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return node["data"][cmd], nil
}

func DeleteNode(args map[string]string) (interface{}, *errors.HccError) {
	// uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid errors} }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	if e := json.Unmarshal(result, &node); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return node["data"][cmd], nil
}

func CreateNodeDetail(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_node_detail"
	query := "mutation _ { " + cmd + arguments + "{ node_uuid cpu_model cpu_processors cpu_threads errors} }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var nodeDetail map[string]map[string]model.NodeDetail
	if e := json.Unmarshal(result, &nodeDetail); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return nodeDetail["data"][cmd], nil
}

func DeleteNodeDetail(args map[string]string) (interface{}, *errors.HccError) {
	// node_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_node_detail"
	query := "mutation _ { " + cmd + arguments + "{ node_uuid errors} }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var nodeDetail map[string]map[string]model.NodeDetail
	if e := json.Unmarshal(result, &nodeDetail); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return nodeDetail["data"][cmd], nil
}
