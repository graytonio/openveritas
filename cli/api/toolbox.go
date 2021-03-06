package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Property struct {
	mgm.DefaultModel `bson:",inline"`
	PropertyName     string             `json:"prop_name" bson:"prop_name"`
	PropertyValue    interface{}        `json:"prop_value" bson:"prop_value"`
	NodeName         string             `json:"node_name" bson:"node_name"`
	NodeId           primitive.ObjectID `json:"node_id" bson:"node_id"`
}

type Node struct {
	mgm.DefaultModel `bson:",inline"`
	NodeName         string `json:"node_name" bson:"node_name"`
}

func appendToHostString(host string, routes ...string) string {
	route := strings.Join(routes, "")
	u, _ := url.Parse(host)
	u.Path = path.Join(u.Path, route)
	return u.String()
}

func parseJSONBody(resp *http.Response, data interface{}) *Error {
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, data)
	if err != nil {
		return createError(err)
	}
	return nil
}

func apiGetRequest(route string) (*http.Response, *Error) {
	resp, err := http.Get(route)
	if err != nil {
		return nil, createError(err)
	}

	if resp.StatusCode != http.StatusOK {
		errResp := handleError(resp)
		return nil, errResp
	}

	return resp, nil
}

func apiPutRequest(route string, body []byte) (*http.Response, *Error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, route, bytes.NewBuffer(body))
	if err != nil {
		return nil, createError(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, createError(err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		errResp := handleError(resp)
		return nil, errResp
	}

	return resp, nil
}

func apiDeleteRequest(route string) (*http.Response, *Error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		return nil, createError(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, createError(err)
	}

	if resp.StatusCode != http.StatusOK {
		errResp := handleError(resp)
		return nil, errResp
	}

	return resp, nil
}

func handleError(resp *http.Response) *Error {
	var errResp Error
	err := parseJSONBody(resp, &errResp)
	if err != nil {
		return err
	}
	return &errResp
}

func createError(err error) *Error {
	return &Error{Code: 500, Message: err.Error()}
}
