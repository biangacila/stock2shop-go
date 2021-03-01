package authentication

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"testing"
)

func TestRegisterNewUser1(t *testing.T) {

	u := domain.User{}
	u.AppName = io.APP_NAME
	u.Username = "biangacila@gmail.com"
	u.Name = "Merveilleux"
	u.Surname = "Biangacila"
	u.Org = "C100002"
	u.Role = "super"
	u.Password = GetMd5("admin")
	u.FullName = "Merveilleux Biangacila"
	u.Email = "biangacila@gmail.com"

	profile := make(map[string]interface{})
	profile["Profession"] = "General Manager"

	u.Profile = profile

	io.InsertTable(io.DB_NAME, "User", u)
	DisplayObject("New USER", u)
}

func TestRegisterNewUser(t *testing.T) {

	u := domain.User{}
	u.AppName = io.APP_NAME
	u.Username = "biangacila@gmail.com"
	u.Name = "Merveilleux"
	u.Surname = "Biangacila"
	u.Org = "C100004"
	u.Role = "super"
	u.Password = GetMd5("admin")
	u.FullName = "Merveilleux Biangacila"
	u.Email = "biangacila@gmail.com"

	profile := make(map[string]interface{})
	profile["Profession"] = "General Manager"

	u.Profile = profile

	io.InsertTable(io.DB_NAME, "User", u)
	DisplayObject("New USER", u)
}
