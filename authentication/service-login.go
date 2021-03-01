package authentication

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"encoding/json"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"github.com/robbert229/jwt"
	"time"
)

type ServiceLogin struct {
	Username string
	Password string
	/* Working variable */
	token    string
	user     domain.User
	hasLogin bool
}

func (obj *ServiceLogin) NewLogin() (bool, domain.User, string, string) {
	//todo let find out if our user exist base on the username
	user, okUser := obj.checkIfUsernameExist()
	if !okUser {
		return false, domain.User{}, "", "Invalid credentials"
	}

	//todo let compare passport
	if user.Password != global.GetMd5(obj.Password) {
		return false, domain.User{}, "", "Invalid credentials"
	}
	/*if user.Password != obj.Password {
		return false, User{}, "", "Invalid credentials"
	}*/
	user.Password = ""
	obj.user = user

	//todo let create our token
	token := obj.createToken()

	//todo let send out our success result

	return true, user, token, ""
}
func (obj *ServiceLogin) createToken() string {
	secret := SignKey
	algorithm := jwt.HmacSha256(string(secret))
	claims := jwt.NewClaim()
	str, _ := json.Marshal(obj.user)
	my := make(map[string]interface{})
	json.Unmarshal(str, &my)
	for key, val := range my {
		value := fmt.Sprintf("%v", val)
		claims.Set(key, value)
	}
	claims.SetTime("exp", time.Now().Add(8*time.Hour))
	dt, hr := global.GetDateAndTimeString()
	claims.Set("date", dt)
	claims.Set("time", hr)
	token, err := algorithm.Encode(claims)
	tokenString := fmt.Sprintf("%s", token)
	if algorithm.Validate(token) != nil {
		panic(err)
	}
	TokensList[tokenString] = obj.user
	return tokenString
}
func (obj *ServiceLogin) checkIfUsernameExist() (domain.User, bool) {
	var ls []domain.User
	qry := fmt.Sprintf("select * from %v.%v where appName='%v' and %v='%v'",
		io.DB_NAME, tableLogin, io.APP_NAME, fieldUsername, obj.Username)
	res, _ := io.RunQueryCass2(qry)
	_ = json.Unmarshal([]byte(res), &ls)

	if len(ls) > 0 {
		return ls[0], true
	}
	return domain.User{}, false
}
