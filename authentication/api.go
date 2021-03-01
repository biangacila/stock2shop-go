package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"net/http"
	"stock2shop-go/domain"
	"strings"
)

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
	User  domain.User
}

func WsUser_New(w http.ResponseWriter, r *http.Request) {
	_, dataString := global.GetPostedDataMapAndString(r)
	var user domain.User
	hub := ServiceUser{}
	json.Unmarshal([]byte(dataString), &user)
	ls := hub.New(user)
	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = hub
	my["RESULT"] = ls
}

func WsUser_List(w http.ResponseWriter, r *http.Request) {
	_, dataString := global.GetPostedDataMapAndString(r)
	hub := ServiceUser{}
	json.Unmarshal([]byte(dataString), &hub)
	ls := hub.List()
	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = hub
	my["RESULT"] = ls
	global.PublishToReact(w, r, my, 200)
}

func WsUser_Find(w http.ResponseWriter, r *http.Request) {
	_, dataString := global.GetPostedDataMapAndString(r)
	hub := ServiceUser{}
	json.Unmarshal([]byte(dataString), &hub)
	ls, _ := hub.FindUser()
	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = hub
	my["RESULT"] = ls
	global.PublishToReact(w, r, my, 200)
}

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	myAuthorization := ""
	for k, v := range r.Header {
		myKey := fmt.Sprintf("%v", k)
		myVal := fmt.Sprintf("%v", v[0])
		if strings.Contains(myKey, "Authorization") {
			arr := strings.Split(myVal, "Bearer ")
			myAuthorization = arr[1]

		}
	}

	//todo check if the token is generate by us
	user, isFind := TokensList[myAuthorization]
	if !isFind {
		w.WriteHeader(http.StatusUnauthorized)
		my := make(map[string]interface{})
		my["error"] = "Unauthorised access to this resource > not find"
		str, _ := json.Marshal(my)
		strJson := string(str)
		fmt.Fprint(w, strJson)
		return
	}

	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = myAuthorization
	my["RESULT"] = user

	JsonResponse(my, w)

}
func getErrorMap(msg string) string {
	my := make(map[string]interface{})
	my["error"] = msg
	str, _ := json.Marshal(my)
	return string(str)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	global.RecoverMe("LoginHandler")
	var serviceLogin ServiceLogin
	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&serviceLogin)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprintf(w, getErrorMap("Error in request"))
		return
	}

	//validate user credentials
	boo, user, tokenString, msg := serviceLogin.NewLogin()
	if !boo {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, getErrorMap(msg))
		return
	}

	fmt.Println("-->OK Request login from: ", serviceLogin.Username, serviceLogin.Password)
	//response := Token{tokenString,user}
	data := make(map[string]interface{})
	data["boo"] = boo
	data["user"] = user
	data["token"] = tokenString
	data["msg"] = msg

	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = serviceLogin
	my["DATA"] = data

	JsonResponse(data, w)
}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	myAuthorization := ""
	for k, v := range r.Header {
		myKey := fmt.Sprintf("%v", k)
		myVal := fmt.Sprintf("%v", v[0])
		if strings.Contains(myKey, "Authorization") {
			arr := strings.Split(myVal, "Bearer ")
			myAuthorization = arr[1]

		}
	}

	//todo check if the token is generate by us
	my := make(map[string]interface{})
	my["error"] = "Unauthorised access to this resource"
	str, _ := json.Marshal(my)
	strJson := string(str)
	_, isFind := TokensList[myAuthorization]
	if !isFind {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, strJson)
		return
	}

	if isVal, err := IsValidToken(myAuthorization); !isVal {
		my["error"] = "Unauthorised access to this resource " + err.Error()
		str, _ := json.Marshal(my)
		strJson = string(str)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, strJson)
		return
	}

	next(w, r)

}

//HELPER FUNCTIONS

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
