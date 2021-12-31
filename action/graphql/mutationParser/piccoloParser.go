package mutationParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

func SignUp(args map[string]string) (interface{}, *errors.HccError) {

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "signup"
	query := `mutation _ { ` + cmd + arguments + `{
        id
		group_id
		authentication
        name
        email
		errors{
			errcode
			errtext
		}
	} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var userData map[string]map[string]model.User
	if e := json.Unmarshal(result, &userData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return userData["data"][cmd], nil
}

func UpdateUser(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_user"
	query := `mutation _ { ` + cmd + arguments + `{
        id
		group_id
		group_name
		authentication
        name
        email
		login_at
		created_at
		errors{
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var userData map[string]map[string]model.User
	if e := json.Unmarshal(result, &userData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return userData["data"][cmd], nil
}

func Unregister(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "unregister"
	query := `mutation _ { ` + cmd + arguments + `{
		id
		errors {
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}
	var userData map[string]map[string]model.User
	if e := json.Unmarshal(result, &userData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return userData["data"][cmd], nil
}

func CreateGroup(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "create_group"
	query := `mutation _ { ` + cmd + arguments + `{
		group_id
        group_name
		errors{
			errcode
			errtext
		}
	} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var groupData map[string]map[string]model.Group
	if e := json.Unmarshal(result, &groupData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return groupData["data"][cmd], nil
}

func UpdateGroup(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "update_group"
	query := `mutation _ { ` + cmd + arguments + `{
		group_id
        group_name
		errors{
			errcode
			errtext
		}
	} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var groupData map[string]map[string]model.Group
	if e := json.Unmarshal(result, &groupData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return groupData["data"][cmd], nil
}

func DeleteGroup(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "delete_group"
	query := `mutation _ { ` + cmd + arguments + `{
		group_id
		errors{
			errcode
			errtext
		}
	} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var groupData map[string]map[string]model.Group
	e := json.Unmarshal(result, &groupData)
	if e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, e.Error())
	}

	return groupData["data"][cmd], nil
}
