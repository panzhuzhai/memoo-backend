package testtool

import (
	"net/http"
	"testing"
)

func TrendingCreators(t *testing.T, init bool) {
	uri := "/api/v1/web-unauthorized/trending-creators?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, NO_AUTH, init, "")
}

func TopCreators(t *testing.T, init bool) {
	uri := "/api/v1/web-unauthorized/top-creators?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, NO_AUTH, init, "")
}
