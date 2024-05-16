package test

import (
	"memoo-backend/initapp"
	"memoo-backend/remoteclient"
	"testing"
)

func TestTwitterVerify(t *testing.T) {
	initapp.Init()
	err := remoteclient.TwitterVerify("https://twitter.com/W8sFgv45Jt16576")
	if err != nil {
		t.Error("Error response:", err)
	}
}
