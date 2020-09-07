package queryParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/lib/errors"
	"hcc/clarinet/model"
)

func Server(args map[string]string) (interface{}, *errors.HccError) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "server"
	query := "query { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid created_at errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string]model.Server
	if e := json.Unmarshal(result, &serverData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverData["data"][cmd], nil
}

func ListServer(args map[string]string) (interface{}, *errors.HccError) {

	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLArgumentError, "Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_server"
	query := "query { " + cmd + arguments + "{ uuid subnet_uuid os server_name server_desc cpu memory disk_size status user_uuid errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string][]model.Server
	if e := json.Unmarshal(result, &serverData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverData["data"][cmd], nil
}

func ServerNode(args map[string]string) (interface{}, *errors.HccError) {
	// uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "server_node"
	query := "query { " + cmd + arguments + "{ uuid server_uuid node_uuid created_at errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	if e := json.Unmarshal(result, &serverNodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverNodeData["data"][cmd], nil
}

func ListServerNode(args map[string]string) (interface{}, *errors.HccError) {
	// server_uuid must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "list_server_node"
	query := "query { " + cmd + arguments + "{ uuid server_uuid node_uuid created_at errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	if e := json.Unmarshal(result, &serverNodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverNodeData["data"][cmd], nil
}

func NumNodesServer(args map[string]string) (interface{}, *errors.HccError) {
	// server_uuid must checked by caller or cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"server_uuid": args["server_uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "num_nodes_server"
	query := "query { " + cmd + arguments + "{ number errors } }"

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var serverNodeNum map[string]map[string]model.ServerNodeNum
	if e := json.Unmarshal(result, &serverNodeNum); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverNodeNum["data"][cmd], nil
}
