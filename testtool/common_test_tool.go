package testtool

import (
	"fmt"
	"io"
	"log"
	"memoo-backend/config"
	"memoo-backend/httpclient"
	"memoo-backend/initapp"
	"memoo-backend/serializer"
	"mime/multipart"
	"os"
	"testing"
)

const (
	AUTH    = "auth"
	NO_AUTH = "no_auth"
)

const (
	APPLICATION_JSON = "application/json"
)

func buildBaseHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Accept"] = "application/json, text/plain, */*"
	headers["Accept-Encoding"] = "gzip, deflate, br, zstd"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9"
	headers["Connection"] = "keep-alive"
	return headers
}

func pullTokenFromLoginResponse(t *testing.T) string {
	resp := Login(t, false)
	if dataMap, ok := resp.Data.(map[string]interface{}); ok {
		if token, ok1 := dataMap["token"].(string); ok1 {
			return token
		}
		return ""
	}
	return ""
}

func buildAuthorizationHeader(t *testing.T) map[string]string {
	token := pullTokenFromLoginResponse(t)
	headers := buildBaseHeaders()
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return headers
}

func assetResult(t *testing.T, uri string, headers map[string]string, bodyByte []byte, method string, auth string, init bool, contentType string) serializer.Response {
	if init {
		initapp.Init()
	}
	switch auth {
	case AUTH:
		headers = buildAuthorizationHeader(t)
	case NO_AUTH:
		headers = buildBaseHeaders()
	default:
		log.Println("no support method")
	}
	if contentType != "" && len(contentType) != 0 {
		headers["Content-Type"] = contentType
	}
	response := serializer.Response{}
	_, err := httpclient.CallRemoteReturnObj(uri, config.AppConfig.TestAttribute.TestApiHost, &response, headers, bodyByte, method)
	if err != nil {
		log.Println(err)
	}
	if response.Code != 200 {
		t.Error("Error response:", response)
	}
	return response
}

func buildFile(writer *multipart.Writer, filePath string, fileName string, fieldName string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
