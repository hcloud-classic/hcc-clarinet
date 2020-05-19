package http

import (
	"encoding/json"
	"errors"
	clarinetData "hcc/clarinet/data"
	"hcc/clarinet/lib/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(moduleName string, needData bool, dataType string, data interface{}, query string) (interface{}, error) {
	var timeout time.Duration
	var url = "http://"
	switch moduleName {
	case "flute":
		timeout = time.Duration(config.Flute.RequestTimeoutMs)
		url += config.Flute.ServerAddress + ":" + strconv.Itoa(int(config.Flute.ServerPort))
		break
	case "harp":
		timeout = time.Duration(config.Harp.RequestTimeoutMs)
		url += config.Harp.ServerAddress + ":" + strconv.Itoa(int(config.Harp.ServerPort))
		break
	case "violin":
		timeout = time.Duration(config.Violin.RequestTimeoutMs)
		url += config.Violin.ServerAddress + ":" + strconv.Itoa(int(config.Violin.ServerPort))
		break
	default:
		return nil, errors.New("unknown module name")
	}
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
			result := string(respBody)

			if strings.Contains(result, "errors") {
				return nil, errors.New(result)
			}

			if needData {
				if data == nil {
					return nil, errors.New("needData marked as true but data is nil")
				}

				switch dataType {
				case "NodeData":
					nodeData := data.(clarinetData.NodeData)
					err = json.Unmarshal([]byte(result), &nodeData)
					if err != nil {
						return nil, err
					}

					return nodeData.Data.Node, nil
				case "AllNodeData":
					allNodeData := data.(clarinetData.AllNodeData)
					err = json.Unmarshal([]byte(result), &allNodeData)
					if err != nil {
						return nil, err
					}

					return allNodeData.Data.AllNode, nil
				case "NumNodeData":
					numNodeData := data.(clarinetData.NumNodeData)
					err = json.Unmarshal([]byte(result), &numNodeData)
					if err != nil {
						return nil, err
					}
				case "NodeDetailData":
					nodeDetailData := data.(clarinetData.NodeDetailData)
					err = json.Unmarshal([]byte(result), &nodeDetailData)
					if err != nil {
						return nil, err
					}

					return nodeDetailData.Data.NodeDetail, nil
				case "OnNodeData":
					onNodeData := data.(clarinetData.OnNodeData)
					err = json.Unmarshal([]byte(result), &onNodeData)
					if err != nil {
						return nil, err
					}

					return onNodeData.Data.Result, nil
				case "CreateNodeData":
					createNodeData := data.(clarinetData.CreateNodeData)
					err = json.Unmarshal([]byte(result), &createNodeData)
					if err != nil {
						return nil, err
					}
					return numNodeData.Data.NumNode, nil

				case "UpdateNodeData":
					updateNodeData := data.(clarinetData.UpdateNodeData)
					err = json.Unmarshal([]byte(result), &updateNodeData)
					if err != nil {
						return nil, err
					}

					return updateNodeData.Data.Node, nil
				case "DeleteNodeData":
					deleteNodeData := data.(clarinetData.DeleteNodeData)
					err = json.Unmarshal([]byte(result), &deleteNodeData)
					if err != nil {
						return nil, err
					}

					return deleteNodeData.Data.Node, nil
				case "CreateNodeDetailData":
					createNodeDetailData := data.(clarinetData.CreateNodeDetailData)
					err = json.Unmarshal([]byte(result), &createNodeDetailData)
					if err != nil {
						return nil, err
					}

					return createNodeDetailData.Data.NodeDetail, nil

				case "DeleteNodeDetailData":
					deleteNodeDetailData := data.(clarinetData.DeleteNodeDetailData)
					err = json.Unmarshal([]byte(result), &deleteNodeDetailData)
					if err != nil {
						return nil, err
					}

					return deleteNodeDetailData.Data.NodeDetail, nil
				case "ServerData":
					serverData := data.(clarinetData.ServerData)
					err = json.Unmarshal([]byte(result), &serverData)
					if err != nil {
						return nil, err
					}

					return serverData.Data.Server, nil
				case "ListServerData":
					listServerData := data.(clarinetData.ListServerData)
					err = json.Unmarshal([]byte(result), &listServerData)
					if err != nil {
						return nil, err
					}

					return listServerData.Data.ListServer, nil

				case "AllServerData":
					allServerData := data.(clarinetData.AllServerData)
					err = json.Unmarshal([]byte(result), &allServerData)
					if err != nil {
						return nil, err
					}

					return allServerData.Data.AllServer, nil
				case "NumServerData":
					numServerData := data.(clarinetData.NumServerData)
					err = json.Unmarshal([]byte(result), &numServerData)
					if err != nil {
						return nil, err
					}

					return numServerData.Data.NumServer, nil
				case "CreateServerData":
					createServerData := data.(clarinetData.CreateServerData)
					err = json.Unmarshal([]byte(result), &createServerData)
					if err != nil {
						return nil, err
					}

				default:
					return nil, errors.New("unknown data type")
				}
			}

			return result, nil
		}

		return nil, err
	}

	return nil, errors.New("http response returned error code")
}
