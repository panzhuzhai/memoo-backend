package testtool

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"
)

func UserCreate(t *testing.T, init bool) {
	writer, body := createUserWriter()
	writer.Close()
	uri := "/api/v1/web-oriented/user"
	assetResult(t, uri, nil, body.Bytes(), http.MethodPost, AUTH, init, writer.FormDataContentType())
}

func createUserWriter() (*multipart.Writer, *bytes.Buffer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("twitter", "333")
	writer.WriteField("userBio", "musk coin")
	writer.WriteField("userName", "yuzhoufeichuan")
	writer.WriteField("websiteUrl", "description")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_20-31-49.png", "image_2024-04-28_20-31-49.png", "profileBanner")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_16-30-15.png", "image_2024-04-28_16-30-15.png", "profileBanner")
	buildFile(writer, "C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-04-28_16-30-15.png", "image_2024-04-28_16-30-15.png", "profileImage")
	return writer, body
}

func UserViewOthersList(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/user-view-others-list?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func UserViewOthers(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/user-view-others"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}
