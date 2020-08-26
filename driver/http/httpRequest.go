package http

import (
	"encoding/json"
	"errors"
	"hcc/clarinet/lib/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(moduleName string, query string) ([]byte, error) {
	var timeout time.Duration
	var url = "http://"

	timeout = time.Duration(config.Piccolo.RequestTimeoutMs)
	url += config.Piccolo.ServerAddress + ":" + strconv.Itoa(int(config.Piccolo.ServerPort))
	url += "/graphql?query=" + queryURLEncoder(query)

	client := &http.Client{Timeout: timeout * time.Millisecond}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Check response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var result map[string]interface{}
			json.Unmarshal(respBody, &result)
			errBody, errChk := result["errors"]

			if !errChk {
				return respBody, nil
			}
			errMsg := ""

			for index, msg := range errBody.([]interface{}) {
				errMsg += strconv.Itoa(index+1) + ":" + msg.(map[string]interface{})["message"].(string) + "\n"
			}
			return nil, errors.New(errMsg)
		}

		return nil, err
	}

	return nil, errors.New("http response returned error code" + strconv.Itoa(resp.StatusCode))
}
