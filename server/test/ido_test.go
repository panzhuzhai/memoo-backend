package token_ranking_test

import (
	"memoo-backend/testtool"
	"testing"
)

func TestIdoActive(t *testing.T) {
	testtool.IdoActive(t, true)
}

func TestIdoCompleted(t *testing.T) {
	testtool.IdoCompleted(t, true)
}

func TestIdoUpcoming(t *testing.T) {
	testtool.IdoUpcoming(t, true)
}
