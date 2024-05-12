package token_ranking_test

import (
	"memoo-backend/testtool"
	"testing"
)

/*
*test TestLogin
 */
func TestLogin(t *testing.T) {
	testtool.Login(t, true)
}

/*
*test TestRefreshToken
 */
func TestRefreshToken(t *testing.T) {
	testtool.RefreshToken(t, true)
}

/*
*test TestLogout
 */
func TestLogout(t *testing.T) {
	testtool.Logout(t, true)
}
