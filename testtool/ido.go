package testtool

import (
	"net/http"
	"testing"
)

func IdoActive(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-active?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func IdoCompleted(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-completed?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func IdoUpcoming(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-upcoming?pageNumber=1&pageSize=4"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}
