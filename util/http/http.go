/*
@Time : 9/5/19
@Author : jzd
@Project: sigmaop
*/
package http

import (
	"bytes"
	"encoding/json"
	"mvc/util/logs"
	"net/http"
)

//https post util
func CreateReqBody(method string, address string, mapBody map[string]interface{}) (*http.Request, error) {
	var reader *bytes.Reader
	if mapBody != nil {
		bytesData, err := json.Marshal(mapBody)
		if err != nil {
			logs.Error("get body err", err)
			return nil, err
		}
		reader = bytes.NewReader(bytesData)
	} else {
		reader = nil
	}
	req, err := http.NewRequest(method, address, reader)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return req, err
}
