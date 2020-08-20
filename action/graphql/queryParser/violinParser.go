package queryParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"
)

func Server(args map[string]string) (interface{}, error) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "server"
	query := "query { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string]model.Server
	err = json.Unmarshal(result, &serverData)
	if err != nil {
		return nil, err
	}
	return serverData["data"][cmd], nil
}

func ListServer(args map[string]string) (interface{}, error) {

	if (args["row"] != "0") != (args["page"] != "0") {

		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_server"
	query := "query { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string][]model.Server
	err = json.Unmarshal(result, &serverData)
	if err != nil {
		return nil, err
	}
	return serverData["data"][cmd], nil
}

func ServerNode(args map[string]string) (interface{}, error) {
	// uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "server_node"
	query := "query { " + cmd + arguments + "{ uuid server_uuid node_uuid created_at } }"

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	err = json.Unmarshal(result, &serverNodeData)
	if err != nil {
		return nil, err
	}
	return serverNodeData["data"][cmd], nil
}

func ListServerNode(args map[string]string) (interface{}, error) {
	// server_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "list_server_node"
	query := "query { " + cmd + arguments + "{ uuid server_uuid node_uuid created_at } }"

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	err = json.Unmarshal(result, &serverNodeData)
	if err != nil {
		return nil, err
	}
	return serverNodeData["data"][cmd], nil
}

func NumNodesServer(args map[string]string) (interface{}, error) {
	// server_uuid must checked by caller or cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "num_nodes_server"
	query := "query { " + cmd + arguments + "{ number } }"

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	var serverNodeNum map[string]map[string]model.ServerNodeNum
	err = json.Unmarshal(result, &serverNodeNum)
	if err != nil {
		return nil, err
	}
	return serverNodeNum["data"][cmd], nil
}
