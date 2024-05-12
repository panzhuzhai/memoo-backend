package token_ranking

import (
	"memoo-backend/testtool"
	"testing"
)

func TestTrendingTokens(t *testing.T) {
	testtool.TrendingTokens(t, true)
}

func TestTopTokens(t *testing.T) {
	testtool.TopTokens(t, true)
}
