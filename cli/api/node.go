package api

import "encoding/json"

func GetAllNodes(host string) ([]Node, error) {
	route := appendToHostString(host, "/node")
	resp, err := apiGetRequest(route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nodes []Node
	err = parseJSONBody(resp, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func GetNodeByName(host string, node_name string) (*Node, error) {
	route := appendToHostString(host, "/node/", node_name)
	resp, err := apiGetRequest(route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var node Node
	err = parseJSONBody(resp, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func PutNode(host string, node_name string, new_name string) (int, error) {
	route := appendToHostString(host, "/node/", node_name)
	node := &Node{
		NodeName: new_name,
	}

	json, err := json.Marshal(node)
	if err != nil {
		return -1, err
	}

	resp, err := apiPutRequest(route, json)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}

func DeleteNode(host string, node_name string) (int, error) {
	route := appendToHostString(host, "/node/", node_name)
	resp, err := apiDeleteRequest(route)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}
