package queryParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

var errQuery string = `errors { errcode errtext }`

func Login(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}
	cmd := "login"
	query := `
	query { ` + cmd + arguments + `{
		token ` +
		errQuery +
		`} }`
	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var loginData map[string]map[string]model.Login
	if e := json.Unmarshal(result, &loginData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return loginData["data"][cmd], nil
}

func CheckToken(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "check_token"
	query := `
	query {` + cmd + arguments + ` {
		isvalid ` +
		errQuery +
		`} }`

	result, err := http.DoHTTPRequest(query)
	if err != nil {
		return nil, err
	}

	var valid map[string]map[string]model.Valid
	if e := json.Unmarshal(result, &valid); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, e.Error())
	}
	return valid["data"][cmd], nil
}
