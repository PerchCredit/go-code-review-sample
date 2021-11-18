package rep

import (
	"code-review-sample/mod"
	"errors"
	"time"
)

type UserAccessor interface {
	A(mod.User) (mod.User, error)
}

type userRepository struct {
	U []*mod.User
	N int
}

func NewRepo() *userRepository {
	return &userRepository{
		make([]*mod.User, 0),
		1,
	}
}

func (ur *userRepository) AddUser(u mod.User) (mod.User, error) {
	if u.Id != 0 {
		return mod.User{}, errors.New("invalid")
	}
	u.Id = ur.N
	u.CT = time.Now().UTC()
	ur.N++
	ur.U = append(ur.U, &u)
	return u, nil
}
