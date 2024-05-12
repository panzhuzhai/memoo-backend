package testtool

import (
	"net/http"
	"testing"
)

func IdoUpcomingDetail(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-upcoming-detail"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func IdoActiveDetail(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-active-detail"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func IdoLaunchedDetail(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-launched-detail"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}

func IdoLaunchedDetailTop10List(t *testing.T, init bool) {
	uri := "/api/v1/web-oriented/ido-launched-detail-top10"
	assetResult(t, uri, nil, nil, http.MethodGet, AUTH, init, "")
}
