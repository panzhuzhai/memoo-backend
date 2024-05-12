package testtool

import (
	"net/http"
	"testing"
)

func TrendingTokens(t *testing.T, init bool) {
	uri := "/api/v1/web-unauthorized/trending-tokens?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, NO_AUTH, init, "")
}

func TopTokens(t *testing.T, init bool) {
	uri := "/api/v1/web-unauthorized/top-tokens?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, NO_AUTH, init, "")
}
