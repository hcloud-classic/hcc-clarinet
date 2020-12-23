package mutationParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
	"hcc/clarinet/model"
)

func CreateServer(args map[string]string) (interface{}, error) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.New("Check flag value of " + ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_server"
	query := "mutation _ { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size user_uuid } }"

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

func UpdateServer(args map[string]string) (interface{}, error) {

	if argumentParser.CheckArgsMin(args, 2, "uuid") {
		return nil, errors.New("Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_server"
	query := "mutation _ { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

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

func DeleteServer(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_server"
	query := "mutation _ { " + cmd + arguments + "{ uuid } }"

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

func CreateServerNode(args map[string]string) (interface{}, error) {
	// serverUUID & nodeUUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
		"node_uuid":   args["node_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "create_server_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid server_uuid node_uuid created_at } }"

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

func DeleteServerNode(args map[string]string) (interface{}, error) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_server_node"
	query := "mutation _ { " + cmd + arguments + "{ uuid } }"

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
