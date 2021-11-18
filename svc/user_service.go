package svc

import (
	"code-review-sample/mod"
	"code-review-sample/rep"
	"errors"
)

type UserServicer interface {
	Create(mod.User) (mod.User, error)
}

type userService struct {
}

func New() *userService {
	return &userService{}
}

func (s *userService) Create(u mod.User) (mod.User, error) {
	if len(u.FN) < 1 || len(u.LN) < 1 || u.LN == "RESTRICTED" {
		return mod.User{}, errors.New("invalid")
	}

	r := rep.NewRepo()
	nu, _ := r.AddUser(u)
	return nu, nil
}
