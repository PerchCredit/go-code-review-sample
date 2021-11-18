package main

import (
	"code-review-sample/mod"
	"code-review-sample/svc"
	"fmt"
)

func main() {
	user := mod.User{
		FN: "John",
		LN: "Doe",
	}

	userService := svc.New()
	newUser, err := userService.Create(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newUser)
	}
}
