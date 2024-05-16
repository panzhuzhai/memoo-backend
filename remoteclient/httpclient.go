package remoteclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CallRemote(url string, ApiHost string, headers map[string]string, bodyByte []byte, method string) ([]byte, error) {
	url = fmt.Sprintf("%s"+url, ApiHost)
	log.Println("call url:", url)
	client := &http.Client{}
	var body io.Reader
	if bodyByte != nil && len(bodyByte) != 0 {
		body = bytes.NewBuffer(bodyByte)
	}
	var responseBytes []byte
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return responseBytes, err
	}
	// Set the request headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return responseBytes, err
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseBytes, err
	}
	log.Println("Response:", string(response))
	return response, nil
}

func CallRemoteReturnObj(url string, ApiHost string, response interface{}, headers map[string]string, bodyByte []byte, method string) (interface{}, error) {
	responseStr, err := CallRemote(url, ApiHost, headers, bodyByte, method)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(responseStr, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
