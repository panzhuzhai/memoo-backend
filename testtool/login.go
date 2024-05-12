package testtool

import (
	"memoo-backend/serializer"
	"net/http"
	"testing"
)

func Login(t *testing.T, init bool) serializer.Response {
	uri := "/api/v1/auth/login"
	byteData := []byte("{\"address\": \"bc1pxl8w2zjafwcywrnnfjtfqvjt9pk75szp2ufqfdrssre55w40fynstngzsd\",\"publicKey\": \"0313cabb6f6912f902d13f5b53ae92e3edb2877c945aa2ac5d7236d683760815a6\",\"message\": \"1715497844536\",\"signature\": \"H0stcSF681jbaGhxTyI0k0SO7QR1iy7m9fXKgyg2uqL6WeQW1YVHzKT2ZNqslzlrQZJlhAbzpE+LKQ1G5XgIU34=\",\"chain\": \"Bitcoin\"}")
	response := assetResult(t, uri, nil, byteData, http.MethodPost, NO_AUTH, init, APPLICATION_JSON)
	return response
}

/*
*test RefreshToken
 */
func RefreshToken(t *testing.T, init bool) {
	uri := "/api/v1/auth/refresh-token"
	assetResult(t, uri, nil, nil, http.MethodPost, AUTH, init, APPLICATION_JSON)
}

/*
*test RefreshToken
 */
func Logout(t *testing.T, init bool) {
	uri := "/api/v1/auth/logout"
	assetResult(t, uri, nil, nil, http.MethodPost, AUTH, init, APPLICATION_JSON)
}
