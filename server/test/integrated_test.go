package token_ranking

import (
	"memoo-backend/initapp"
	"memoo-backend/testtool"
	"testing"
)

func TestIntegrated(t *testing.T) {
	initapp.Init()
	/***********login***********/
	testtool.Login(t, false)
	testtool.RefreshToken(t, false)
	testtool.Logout(t, false)

	/***********Creator Ranking***********/
	testtool.TrendingCreators(t, false)
	testtool.TopCreators(t, false)

	/***********token Ranking***********/
	testtool.TrendingTokens(t, false)
	testtool.TopTokens(t, false)

	/***********user***********/
	testtool.UserCreate(t, false)
	testtool.UserViewOthersList(t, false)
	testtool.UserViewOthers(t, false)

	/***********ido***********/
	testtool.IdoActive(t, false)
	testtool.IdoCompleted(t, false)
	testtool.IdoUpcoming(t, false)

	/***********ido-detail***********/
	testtool.IdoUpcomingDetail(t, false)
	testtool.IdoActiveDetail(t, false)
	testtool.IdoLaunchedDetail(t, false)
	testtool.IdoLaunchedDetailTop10List(t, false)

	/***********project***********/
	testtool.ProjectCreate(t, false)

	/***********token***********/
	testtool.TokenCreate(t, false)
	testtool.TokenList(t, false)

}
