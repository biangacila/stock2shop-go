package authentication

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"testing"
)

func TestServiceUser_List(t *testing.T) {
	hub := ServiceUser{}
	//hub.Org = "easipath"
	ls := hub.List()
	global.DisplayObject("Users", ls)
}
func TestServiceUser_UpdateField(t *testing.T) {
	hub := ServiceUser{}
	hub.Username = "kim.scott@sivuyileservices.co.za"
	hub.FieldName = "password"
	hub.FieldValue = "sivuadmin"
	hub.UpdateField()
}
func TestServiceUser_FindUser(t *testing.T) {
	//io.SayHello()
	hub := ServiceUser{}
	hub.Username = "biangacila@gmail.com"
	find, boo := hub.FindUser()
	if !boo {
		fmt.Println("--> sorry user not find!")

	}
	global.DisplayObject("User", find)
}
func TestServiceUser_New(t *testing.T) {
	io.SayHello()
	user := domain.User{
		Email:    "peterd@marginmentor.co.za",
		Name:     "Peter",
		Surname:  "David",
		Phone:    "0729139504",
		Org:      "C100003",
		Role:     "super",
		Password: "admin",
	}

	hub := ServiceUser{}
	hub.New(user)
}

// update pmis.User set password='8be0f0b5265d4f5b8d0a126f0c18230f' where appname='pmis' and username='kim.scott@sivuyileservices.co.za';
