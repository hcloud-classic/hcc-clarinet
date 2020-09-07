package queryParser

import (
	"encoding/json"

	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/lib/errors"
)

func Login(args map[string]string) (interface{}, *errors.HccError) {
	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "login"
	query := "query { " + cmd + arguments + "{ token errors } }"
	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var loginData map[string]map[string]string
	if e := json.Unmarshal(result, &loginData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return loginData["data"][cmd], nil
}
