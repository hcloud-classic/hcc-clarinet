package queryParser

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/http"
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
	query := "query { server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at } }"

	var serverData struct {
		Data struct {
			Server model.Server `json:"server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &serverData)
	if err != nil {
		return nil, err
	}
	return serverData.Data.Server, nil
}

func ListServer(args map[string]string) (interface{}, error) {

	if (args["row"] != "0") != (args["page"] != "0") {

		return nil, errors.New("Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}
	query := "query { list_server(" + arguments + ") { uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid } }"

	var listServerData struct {
		Data struct {
			ListServer []model.Server `json:"list_server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &listServerData)
	if err != nil {
		return nil, err
	}

	return listServerData.Data.ListServer, nil
}

func ServerNode(args map[string]string) (interface{}, error) {
	// uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}
	query := "query { server_node(" + arguments + ") { uuid server_uuid node_uuid created_at } }"

	var serverNodeData struct {
		Data struct {
			ServerNode model.ServerNode `json:"server_node"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &serverNodeData)
	if err != nil {
		return nil, err
	}
	return serverNodeData.Data.ServerNode, nil
}

func ListServerNode(args map[string]string) (interface{}, error) {
	// server_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}
	query := "query { list_server_node(" + arguments + ") { uuid server_uuid node_uuid created_at } }"

	var listServerNodeData struct {
		Data struct {
			ListServerNode []model.ServerNode `json:"list_server_node"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &listServerNodeData)
	if err != nil {
		return nil, err
	}
	return listServerNodeData.Data.ListServerNode, nil
}

func AllServerNode() (interface{}, error) {
	query := "query { all_server_node { uuid server_uuid node_uuid created_at } }"

	var allServerNodeData struct {
		Data struct {
			AllServerNode []model.ServerNode `json:"all_server_node"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &allServerNodeData)
	if err != nil {
		return nil, err
	}
	return allServerNodeData.Data.AllServerNode, nil

}

func NumNodesServer(args map[string]string) (interface{}, error) {
	// server_uuid must checked by caller or cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}
	query := "query { num_nodes_server(" + arguments + ") { number } }"

	var numNodesServer struct {
		Data struct {
			NumNodesServer model.ServerNodeNum `json:"num_nodes_server"`
		} `json:"data"`
	}
	result, err := http.DoHTTPRequest("violin", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &numNodesServer)
	if err != nil {
		return nil, err
	}
	return numNodesServer.Data.NumNodesServer, nil
}
