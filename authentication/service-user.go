package authentication

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"encoding/json"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"strings"
)

type ServiceUser struct {
	Org        string
	Username   string
	FieldName  string
	FieldValue string
}

func (obj *ServiceUser) UpdateField() {
	if strings.ToLower(obj.FieldName) == "password" {
		obj.FieldValue = global.GetMd5(obj.FieldValue)
	}
	var table = "User"
	qry := fmt.Sprintf("update %v.%v set %v='%v' where username='%v' ",
		io.DB_NAME, table, obj.FieldName, obj.FieldValue, obj.Username)
	_, _ = io.RunQueryCass2(qry)
}
func (obj *ServiceUser) New(in domain.User) string {
	in.Password = global.GetMd5(in.Password)
	in.Email = strings.ToLower(in.Email)
	in.Email = strings.Trim(in.Email, "")
	in.Email = strings.Replace(in.Email, "%", "", 1000)

	in.Username = in.Email
	in.AppName = io.AppName
	io.InsertTable(io.DB_NAME, "User", in)

	return "OK"
}
func (obj *ServiceUser) New2(in domain.User) string {
	in.Password = global.GetMd5(in.Password)
	in.Status = "trial"
	in.FullName = in.Name + " " + in.Surname
	in.FullName = strings.Title(strings.ToLower(in.FullName))
	in.Email = strings.ToLower(in.Email)
	in.Email = strings.Trim(in.Email, "")
	in.Email = strings.Replace(in.Email, "%", "", 1000)

	in.Username = in.Email

	var table = "User"
	io.LibCassInsertQuery(io.DB_NAME, []string{table}, in)
	return "OK"
}
func (obj *ServiceUser) FindUser() (domain.User, bool) {
	var table = "User"
	var ls []domain.User
	qry := fmt.Sprintf("select * from %v.%v where appName='%v' and username='%v'",
		io.DB_NAME, table, io.APP_NAME, obj.Username)
	res, _ := io.RunQueryCass2(qry)
	_ = json.Unmarshal([]byte(res), &ls)

	if len(ls) > 0 {
		return ls[0], true
	}
	return domain.User{}, false
}
func (obj *ServiceUser) List() []domain.User {
	var table = "User"
	var ls []domain.User
	qry := fmt.Sprintf("select * from %v.%v where appName='%v' ",
		io.DB_NAME, table, io.APP_NAME)
	res, _ := io.RunQueryCass2(qry)
	_ = json.Unmarshal([]byte(res), &ls)
	var out []domain.User
	for _, row := range ls {
		row.Password = ""
		if obj.Org == "" {
			out = append(out, row)
		} else {
			if row.Org == obj.Org {

				out = append(out, row)
			}
		}

	}
	return out
}
