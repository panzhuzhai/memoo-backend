package testtool

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"
)

func TokenCreate(t *testing.T, init bool) {
	writer, body := createTokenWriter()
	writer.Close()
	uri := "/api/v1/web-oriented/token"
	assetResult(t, uri, nil, body.Bytes(), http.MethodPost, AUTH, init, writer.FormDataContentType())
}

func createTokenWriter() (*multipart.Writer, *bytes.Buffer) {
	writer, body := createProjectWriter()
	writer.WriteField("tokenName", "musk")
	writer.WriteField("contractAddress", "bc1pxl8w2zjafwcywrnnfjtfqvjt9pk75szp2ufqfdrssre55w40fynstngzsd")
	writer.WriteField("lPContractAddress", "0x32C21A937452Fb191e0209A5FCe20F071B9CB569")
	writer.WriteField("tokenDescription", "tokenDescription")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_10-14-05.png", "image_2024-04-28_10-14-05.png", "tokenIcon")
	return writer, body
}

func TokenList(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/token?pageNumber=1&pageSize=4&status"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}
