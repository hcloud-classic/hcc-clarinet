package mutationParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

const errQuery string = " errors { errcode errtext }"

func CreateServer(args map[string]string) (interface{}, *errors.HccError) {
	if b, ef := argumentParser.CheckArgsAll(args, len(args)); b {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Check flag value of "+ef)
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_server"
	query := `mutation _ { ` + cmd + arguments + `{
		uuid
		server_name ` +
		errQuery +
		`} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string]model.Server
	if e := json.Unmarshal(result, &serverData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverData["data"][cmd], nil
}

func UpdateServer(args map[string]string) (interface{}, *errors.HccError) {

	if argumentParser.CheckArgsMin(args, 2, "uuid") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLParsingError, "Need at least 1 more flag except uuid")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_server"
	query := `mutation _ { ` + cmd + arguments + `{
		uuid
		server_name` +
		errQuery +
		`} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string]model.Server
	if e := json.Unmarshal(result, &serverData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverData["data"][cmd], nil
}

func DeleteServer(args map[string]string) (interface{}, *errors.HccError) {
	// UUID flag must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_server"
	query := `mutation _ { ` + cmd + arguments + `{ uuid ` + errQuery + ` } }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var serverData map[string]map[string]model.Server
	if e := json.Unmarshal(result, &serverData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverData["data"][cmd], nil
}

func CreateServerNode(args map[string]string) (interface{}, *errors.HccError) {
	// serverUUID & nodeUUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_server_node"
	query := `mutation _ { ` + cmd + arguments + `{
		uuid
		server_uuid
		node_uuid
		created_at ` +
		errQuery +
		`} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	if e := json.Unmarshal(result, &serverNodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverNodeData["data"][cmd], nil
}

func DeleteServerNode(args map[string]string) (interface{}, *errors.HccError) {
	// UUID must checked by cobra
	arguments, err := argumentParser.GetArgumentStr(map[string]string{
		"uuid": args["uuid"],
	})
	if err != nil {
		return nil, err
	}

	cmd := "delete_server_node"
	query := `mutation _ { ` + cmd + arguments + `{ uuid ` + errQuery + `} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var serverNodeData map[string]map[string]model.ServerNode
	if e := json.Unmarshal(result, &serverNodeData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return serverNodeData["data"][cmd], nil
}
