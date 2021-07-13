package queryParser

import (
	"encoding/json"
	"hcc/clarinet/action/graphql"
	"hcc/clarinet/driver/http"
	"hcc/clarinet/model"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

func ListUser(args map[string]string) (interface{}, *errors.HccError) {
	if (args["row"] != "0") != (args["page"] != "0") {
		return nil, errors.NewHccError(errors.ClarinetGraphQLArgumentError, "Need [BOTH | NEITHER] row & page")
	}

	arguments, err := argumentParser.GetArgumentStr(args)
	if err != nil {
		return nil, err
	}

	cmd := "list_user"
	query := `query { ` + cmd + arguments + `{ 
		user_list {
			id
			group_id
			group_name
			authentication
			name
			email
			login_at
			created_at
		}
		errors{
			errcode
			errtext
		}
	} }`

	result, err := http.DoHTTPRequest("piccolo", query)
	if err != nil {
		return nil, err
	}

	var userListData map[string]map[string]model.UserList
	if e := json.Unmarshal(result, &userListData); e != nil {
		return nil, errors.NewHccError(errors.ClarinetGraphQLJsonUnmarshalError, err.Error())
	}
	return userListData["data"][cmd], nil
}
