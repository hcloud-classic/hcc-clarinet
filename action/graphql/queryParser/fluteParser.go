package queryParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"
)

func Node(args map[string]string) (interface{}, error) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "node"
	query := "query { " + cmd + arguments + "{ uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string]model.Node
	err = json.Unmarshal(result, &nodeData)
	if err != nil {
		return nil, err
	}
	return nodeData["data"][cmd], nil
}

func ListNode(args map[string]string) (interface{}, error) {

	if (args["row"] != "0") != (args["page"] != "0") {

		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_node"
	query := "query { " + cmd + arguments + "{ uuid server_uuid bmc_mac_addr bmc_ip pxe_mac_addr status cpu_cores memory description created_at active } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeData map[string]map[string][]model.Node
	err = json.Unmarshal(result, &nodeData)
	if err != nil {
		return nil, err
	}
	return nodeData["data"][cmd], nil
}

func NumNode() (interface{}, error) {
	cmd := "num_node"
	query := "query { " + cmd + " { number } }"

	result, err := http.DoHTTPRequest("flute", query)
	if err != nil {
		return nil, err
	}

	var nodeNum map[string]map[string]model.NodeNum
	err = json.Unmarshal(result, &nodeNum)
	if err != nil {
		return nil, err
	}
	return nodeNum["data"][cmd], nil
}

func NodeDetail(args map[string]string) (interface{}, error) {
	// node_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"node_uuid": args["node_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "detail_node"
	query := "query { " + cmd + arguments + "{ node_uuid cpu_model cpu_processors cpu_threads } }"

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
