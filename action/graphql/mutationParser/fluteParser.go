package mutationParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
	"hcc/clarinet/model"
)

// Power on Node
func OnOffNode(args map[string]string, state model.PowerState) (interface{}, error) {
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
		return nil, errors.New("Undefined Power state")
	}

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string]string
	err = json.Unmarshal(result, &nodeData)
	if err != nil {
		return nil, err
	}
	return nodeData["data"][cmd], nil
}

func CreateNode(args map[string]string) (interface{}, error) {
	// bmc_ip & description must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	err = json.Unmarshal(result, &node)
	if err != nil {
		return nil, err
	}
	return node["data"][cmd], nil
}

func UpdateNode(args map[string]string) (interface{}, error) {
	// bmc_ip must checked by cobra
	if argumentParser.CheckArgsMin(args, 2, "bmc_ip") {
		return nil, errors.New("Need at least 1 more flag except bmc_ip")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid bmc_mac_addr } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	err = json.Unmarshal(result, &node)
	if err != nil {
		return nil, err
	}
	return node["data"][cmd], nil
}

func DeleteNode(args map[string]string) (interface{}, error) {
	// uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var node map[string]map[string]model.Node
	err = json.Unmarshal(result, &node)
	if err != nil {
		return nil, err
	}
	return node["data"][cmd], nil
}

func CreateNodeDetail(args map[string]string) (interface{}, error) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.New("Check flag value of " + ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_node_detail"
	query := "mutation _ { " + cmd + arguments + "{ node_uuid cpu_model cpu_processors cpu_threads } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeDetail map[string]map[string]model.NodeDetail
	err = json.Unmarshal(result, &nodeDetail)
	if err != nil {
		return nil, err
	}
	return nodeDetail["data"][cmd], nil
}

func DeleteNodeDetail(args map[string]string) (interface{}, error) {
	// node_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_node_detail"
	query := "mutation _ { " + cmd + arguments + "{ node_uuid } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeDetail map[string]map[string]model.NodeDetail
	err = json.Unmarshal(result, &nodeDetail)
	if err != nil {
		return nil, err
	}
	return nodeDetail["data"][cmd], nil
}
