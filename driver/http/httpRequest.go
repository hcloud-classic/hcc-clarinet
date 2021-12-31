package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"hcc/clarinet/lib/config"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(query string) ([]byte, *errors.HccError) {
	var timeout time.Duration
	var url = "http://"

	timeout = time.Duration(config.Piccolo.RequestTimeoutMs)
	url += config.Piccolo.ServerAddress + ":" + strconv.Itoa(int(config.Piccolo.ServerPort))
	url += "/graphql?query=" + queryURLEncoder(query)

	client := &http.Client{Timeout: timeout * time.Millisecond}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.NewHccError(errors.ClarinetDriverRequestError, err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.NewHccError(errors.ClarinetDriverRequestError, err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Check response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var result map[string]interface{}
			err = json.Unmarshal(respBody, &result)
			if err != nil {
				return nil, errors.NewHccError(errors.ClarinetDriverJsonUnmarshalError, err.Error())
			}
			errBody, errChk := result["errors"]

			if !errChk {
				return respBody, nil
			}
			errMsg := ""

			for index, msg := range errBody.([]interface{}) {
				errMsg += strconv.Itoa(index+1) + ":" + msg.(map[string]interface{})["message"].(string) + "\n"
			}
			return nil, errors.NewHccError(errors.ClarinetDriverReceiveError, "DoHTTPReq: "+errMsg)
		}

		return nil, errors.NewHccError(errors.ClarinetDriverParsingError, "DoHTTPReq: "+err.Error())
	}

	return nil, errors.NewHccError(errors.ClarinetDriverResponseError, "http response returned error code"+strconv.Itoa(resp.StatusCode))
}
