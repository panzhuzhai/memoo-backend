package token_ranking

import (
	"memoo-backend/testtool"
	"testing"
)

func TestTrendingCreators(t *testing.T) {
	testtool.TrendingCreators(t, true)
}

func TestTopCreators(t *testing.T) {
	testtool.TopCreators(t, true)
}
