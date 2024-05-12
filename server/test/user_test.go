package token_ranking_test

import (
	"memoo-backend/testtool"
	"testing"
)

/*
*test TestTokenCreate
 */
func TestUserCreate(t *testing.T) {
	testtool.UserCreate(t, true)
}

/*
*test TestTokenCreate
 */
func TestUserViewOthersList(t *testing.T) {
	testtool.UserViewOthersList(t, true)
}

/*
*test TestTokenCreate
 */
func TestUserViewOthers(t *testing.T) {
	testtool.UserViewOthers(t, true)
}
