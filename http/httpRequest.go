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

					return numNodeData.Data.NumNode, nil
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
				case "OffNodeData":
					offNodeData := data.(clarinetData.OffNodeData)
					err = json.Unmarshal([]byte(result), &offNodeData)
					if err != nil {
						return nil, err
					}

					return offNodeData.Data.Result, nil
				case "RestartNodeData":
					restartNodeData := data.(clarinetData.RestartNodeData)
					err = json.Unmarshal([]byte(result), &restartNodeData)
					if err != nil {
						return nil, err
					}

					return restartNodeData.Data.Result, nil
				case "CreateNodeData":
					createNodeData := data.(clarinetData.CreateNodeData)
					err = json.Unmarshal([]byte(result), &createNodeData)
					if err != nil {
						return nil, err
					}

					return createNodeData.Data.Node, nil
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
				case "CreateServerData":
					createServerData := data.(clarinetData.CreateServerData)
					err = json.Unmarshal([]byte(result), &createServerData)
					if err != nil {
						return nil, err
					}

					return createServerData.Data.Server, nil
				case "UpdateServerData":
					updateServerData := data.(clarinetData.UpdateServerData)
					err = json.Unmarshal([]byte(result), &updateServerData)
					if err != nil {
						return nil, err
					}

					return updateServerData.Data.Server, nil
				case "DeleteServerData":
					deleteServerData := data.(clarinetData.DeleteServerData)
					err = json.Unmarshal([]byte(result), &deleteServerData)
					if err != nil {
						return nil, err
					}

					return deleteServerData.Data.Server, nil
				case "CreateServerNodeData":
					createServerNodeData := data.(clarinetData.CreateServerNodeData)
					err = json.Unmarshal([]byte(result), &createServerNodeData)
					if err != nil {
						return nil, err
					}

					return createServerNodeData.Data.Server, nil
				case "DeleteServerNodeData":
					deleteServerNodeData := data.(clarinetData.DeleteServerNodeData)
					err = json.Unmarshal([]byte(result), &deleteServerNodeData)
					if err != nil {
						return nil, err
					}

					return deleteServerNodeData.Data.Server, nil
				case "ServerNodeData":
					serverNodeData := data.(clarinetData.ServerNodeData)
					err = json.Unmarshal([]byte(result), &serverNodeData)
					if err != nil {
						return nil, err
					}

					return serverNodeData.Data.ServerNode, nil
				case "ListServerNodeData":
					listServerNodeData := data.(clarinetData.ListServerNodeData)
					err = json.Unmarshal([]byte(result), &listServerNodeData)
					if err != nil {
						return nil, err
					}

					return listServerNodeData.Data.ListServerNode, nil
				case "AllServerNodeData":
					allServerNodeData := data.(clarinetData.AllServerNodeData)
					err = json.Unmarshal([]byte(result), &allServerNodeData)
					if err != nil {
						return nil, err
					}

					return allServerNodeData.Data.AllServerNode, nil
				case "NumNodesServerData":
					numNodesServerData := data.(clarinetData.NumNodesServerData)
					err = json.Unmarshal([]byte(result), &numNodesServerData)
					if err != nil {
						return nil, err
					}

					return numNodesServerData.Data.NumNodesServer, nil
				case "SubnetData":
					subnetData := data.(clarinetData.SubnetData)
					err = json.Unmarshal([]byte(result), &subnetData)
					if err != nil {
						return nil, err
					}

					return subnetData.Data.Subnet, nil
				case "ListSubnetData":
					listSubnetData := data.(clarinetData.ListSubnetData)
					err = json.Unmarshal([]byte(result), &listSubnetData)
					if err != nil {
						return nil, err
					}

					return listSubnetData.Data.ListSubnet, nil
				case "AllSubnetData":
					allSubnetData := data.(clarinetData.AllSubnetData)
					err = json.Unmarshal([]byte(result), &allSubnetData)
					if err != nil {
						return nil, err
					}

					return allSubnetData.Data.AllSubnet, nil
				case "NumSubnetData":
					numSubnetData := data.(clarinetData.NumSubnetData)
					err = json.Unmarshal([]byte(result), &numSubnetData)
					if err != nil {
						return nil, err
					}

					return numSubnetData.Data.NumSubnet, nil
				case "CreateSubnetData":
					createSubnetData := data.(clarinetData.CreateSubnetData)
					err = json.Unmarshal([]byte(result), &createSubnetData)
					if err != nil {
						return nil, err
					}

					return createSubnetData.Data.Subnet, nil
				case "UpdateSubnetData":
					updateSubnetData := data.(clarinetData.UpdateSubnetData)
					err = json.Unmarshal([]byte(result), &updateSubnetData)
					if err != nil {
						return nil, err
					}

					return updateSubnetData.Data.Subnet, nil
				case "DeleteSubnetData":
					deleteSubnetData := data.(clarinetData.DeleteSubnetData)
					err = json.Unmarshal([]byte(result), &deleteSubnetData)
					if err != nil {
						return nil, err
					}

					return deleteSubnetData.Data.Subnet, nil
				case "CreateDHCPDConfData":
					createDHCPDConfData := data.(clarinetData.CreateDHCPDConfData)
					err = json.Unmarshal([]byte(result), &createDHCPDConfData)
					if err != nil {
						return nil, err
					}

					return createDHCPDConfData.Data.Result, nil
				case "AdaptiveIPData":
					adaptiveIPData := data.(clarinetData.AdaptiveIPData)
					err = json.Unmarshal([]byte(result), &adaptiveIPData)
					if err != nil {
						return nil, err
					}

					return adaptiveIPData.Data.AdaptiveIP, nil
				case "ListAdaptiveIPData":
					listAdaptiveIPData := data.(clarinetData.ListAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &listAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return listAdaptiveIPData.Data.ListAdaptiveIP, nil
				case "AllAdaptiveIPData":
					allAdaptiveIPData := data.(clarinetData.AllAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &allAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return allAdaptiveIPData.Data.AllAdaptiveIP, nil
				case "NumAdaptiveIPData":
					numAdaptiveIPData := data.(clarinetData.NumAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &numAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return numAdaptiveIPData.Data.NumAdaptiveIP, nil
				case "CreateAdaptiveIPData":
					createAdaptiveIPData := data.(clarinetData.CreateAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &createAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return createAdaptiveIPData.Data.AdaptiveIP, nil
				case "UpdateAdaptiveIPData":
					updateAdaptiveIPData := data.(clarinetData.UpdateAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &updateAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return updateAdaptiveIPData.Data.AdaptiveIP, nil
				case "DeleteAdaptiveIPData":
					deleteAdaptiveIPData := data.(clarinetData.DeleteAdaptiveIPData)
					err = json.Unmarshal([]byte(result), &deleteAdaptiveIPData)
					if err != nil {
						return nil, err
					}

					return deleteAdaptiveIPData.Data.AdaptiveIP, nil
				case "AdaptiveIPServerData":
					adaptiveIPServerData := data.(clarinetData.AdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &adaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return adaptiveIPServerData.Data.AdaptiveIPServer, nil
				case "ListAdaptiveIPServerData":
					listAdaptiveIPServerData := data.(clarinetData.ListAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &listAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return listAdaptiveIPServerData.Data.ListAdaptiveIPServer, nil
				case "AllAdaptiveIPServerData":
					allAdaptiveIPServerData := data.(clarinetData.AllAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &allAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return allAdaptiveIPServerData.Data.AllAdaptiveIPServer, nil
				case "NumAdaptiveIPServerData":
					numAdaptiveIPServerData := data.(clarinetData.NumAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &numAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return numAdaptiveIPServerData.Data.NumAdaptiveIPServer, nil
				case "CreateAdaptiveIPServerData":
					createAdaptiveIPServerData := data.(clarinetData.CreateAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &createAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return createAdaptiveIPServerData.Data.AdaptiveIPServer, nil
				case "DeleteAdaptiveIPServerData":
					deleteAdaptiveIPServerData := data.(clarinetData.DeleteAdaptiveIPServerData)
					err = json.Unmarshal([]byte(result), &deleteAdaptiveIPServerData)
					if err != nil {
						return nil, err
					}

					return deleteAdaptiveIPServerData.Data.AdaptiveIPServer, nil
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
