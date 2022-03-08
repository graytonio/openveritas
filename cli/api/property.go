package api

import "encoding/json"

func GetAllPropertiesOfNode(host string, node_name string) ([]Property, error) {
	route := appendToHostString(host, "/node/", node_name, "/prop")
	resp, err := apiGetRequest(route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var props []Property
	err = parseJSONBody(resp, &props)
	if err != nil {
		return nil, err
	}

	return props, nil
}

func GetPropertyOfNodeByName(host string, node_name string, prop_name string) (*Property, error) {
	route := appendToHostString(host, "/node/", node_name, "/prop/", prop_name)
	resp, err := apiGetRequest(route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var prop Property
	err = parseJSONBody(resp, &prop)
	if err != nil {
		return nil, err
	}

	return &prop, nil
}

func PutProp(host string, node_name, prop_name string, prop_value interface{}) (int, error) {
	route := appendToHostString(host, "/node/", node_name, "/prop/", prop_name)
	prop := &Property{
		PropertyName:  prop_name,
		PropertyValue: prop_value,
	}

	json, err := json.Marshal(prop)
	if err != nil {
		return -1, err
	}

	resp, err := apiPutRequest(route, json)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}

func DeleteProp(host string, node_name string, prop_name string) (int, error) {
	route := appendToHostString(host, "/node/", node_name, "/prop/", prop_name)
	resp, err := apiDeleteRequest(route)
	if err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}
