package testtool

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"
)

func ProjectCreate(t *testing.T, init bool) {
	writer, body := createProjectWriter()
	writer.Close()
	uri := "/api/v1/web-oriented/project"
	assetResult(t, uri, nil, body.Bytes(), http.MethodPost, AUTH, init, writer.FormDataContentType())
}

func createProjectWriter() (*multipart.Writer, *bytes.Buffer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("twitter", "333")
	writer.WriteField("ticker", "musk coin")
	writer.WriteField("projectName", "yuzhoufeichuan")
	writer.WriteField("description", "description")
	writer.WriteField("otherLinkStr", "[{\"otherLinkUrl\": \"John\",\"otherLinkUrlType\":\"website\"}, {\"otherLinkUrl\": \"Alice\",\"otherLinkUrlType\":\"tt\"}]")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_20-31-49.png", "image_2024-04-28_20-31-49.png", "banner")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_16-30-15.png", "image_2024-04-28_16-30-15.png", "banner")
	return writer, body
}
