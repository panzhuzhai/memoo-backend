package token_ranking_test

import (
	"memoo-backend/testtool"
	"testing"
)

/*
*test TestTokenCreate
 */
func TestTokenCreate(t *testing.T) {
	testtool.TokenCreate(t, true)
}

func TestTokenList(t *testing.T) {
	testtool.TokenList(t, true)
}
