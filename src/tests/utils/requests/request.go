package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func request(method, route string, body interface{}) *http.Request {
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		panic(fmt.Sprintf("Invalid json request on test: %v %v", method, route))
	}
	req := httptest.NewRequest(method, route, bytes.NewBuffer(jsonResponse))
	req.Header.Add("Content-Type", "application/json")
	return req
}

func Get(route string) *http.Request {
	return request("GET", route, nil)
}

func Post(route string, body interface{}) *http.Request {
	return request("POST", route, body)
}

func Put(route string, body interface{}) *http.Request {
	return request("PUT", route, body)
}

func Delete(route string, body interface{}) *http.Request {
	return request("DELETE", route, body)
}
