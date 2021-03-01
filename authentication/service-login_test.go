package authentication

import (
	"fmt"
	"testing"
)

func TestServiceLogin_NewLogin(t *testing.T) {
	hub := ServiceLogin{}
	hub.Username = "biangacila@gmail.com"
	hub.Password = "admin"
	boo, user, token, msg := hub.NewLogin()
	fmt.Println(":> Has login: ", boo)
	fmt.Println(":> User: ", user)
	fmt.Println(":> token: ", token)
	fmt.Println(":> msg: ", msg)
}
