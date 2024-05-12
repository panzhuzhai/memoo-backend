package token_ranking_test

import (
	"memoo-backend/testtool"
	"testing"
)

func TestIdoUpcomingDetail(t *testing.T) {
	testtool.IdoUpcomingDetail(t, true)
}

func TestIdoActiveDetail(t *testing.T) {
	testtool.IdoActiveDetail(t, true)
}

func TestIdoLaunchedDetail(t *testing.T) {
	testtool.IdoLaunchedDetail(t, true)
}
func TestIdoLaunchedDetailTop10List(t *testing.T) {
	testtool.IdoLaunchedDetailTop10List(t, true)
}
