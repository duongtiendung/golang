package simpleservice

import "testing"

func TestCheckUserExist(t *testing.T) {
	user := User{
		Name:     "Duong Dung",
		Email:    "dungdt@lifetimetech.vn",
		UserName: "dungdt",
	}

	err := RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}
