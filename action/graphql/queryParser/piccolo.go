package queryParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
)

func Login(args map[string]string) (interface{}, error) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "login"
	query := "query { " + cmd + arguments + "{ token } }"
	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var loginData map[string]map[string]string
	err = json.Unmarshal(result, &loginData)
	if err != nil {
		return nil, err
	}
	return loginData["data"][cmd], nil
}
