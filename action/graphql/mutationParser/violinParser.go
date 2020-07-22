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

	query := "mutation _ { create_server (" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size user_uuid } }"

	var createServerData struct {
		Data struct {
			Server model.Server `json:"create_server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &createServerData)
	if err != nil {
		return nil, err
	}
	return createServerData.Data.Server, nil
}

func UpdateServer(args map[string]string) (interface{}, error) {

	if argumentParser.CheckArgsMin(args, 2) {
		return nil, errors.New("Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}
	query := "mutation _ { update_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	var updateServerData struct {
		Data struct {
			Server model.Server `json:"update_server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &updateServerData)
	if err != nil {
		return nil, err
	}
	return updateServerData.Data.Server, nil

}

func DeleteServer(args map[string]string) (interface{}, error) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}
	query := "mutation _ { delete_server(" + arguments + ", status:\"Deleted\") { uuid } }"

	var deleteServerData struct {
		Data struct {
			Server model.Server `json:"delete_server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &deleteServerData)
	if err != nil {
		return nil, err
	}
	return deleteServerData.Data.Server, nil
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
	query := "mutation _ { create_server_node(" + arguments + "\") { uuid server_uuid node_uuid created_at } }"

	var createServerNodeData struct {
		Data struct {
			Server model.ServerNode `json:"create_server_node"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &createServerNodeData)
	if err != nil {
		return nil, err
	}
	return createServerNodeData.Data.Server, nil
}

func DeleteServerNode(args map[string]string) (interface{}, error) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}
	query := "mutation _ { delete_server_node(" + arguments + ") { uuid } }"

	var deleteServerNodeData struct {
		Data struct {
			Server model.ServerNode `json:"delete_server_node"`
		} `json:"data"`
	}

	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &deleteServerNodeData)
	if err != nil {
		return nil, err
	}
	return deleteServerNodeData.Data.Server, nil
}
