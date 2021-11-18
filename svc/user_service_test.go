package svc_test

import (
	"code-review-sample/mod"
	"code-review-sample/svc"
	"testing"
)

func TestCreate(t *testing.T) {
	u := mod.User{
		FN: "Bob",
		LN: "Smith",
	}

	us := svc.New()

	_, e := us.Create(u)
	if e != nil {
		t.Errorf("Got error, expected success")
	}
}
