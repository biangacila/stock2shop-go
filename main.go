package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"stock2shop-go/api"
	_ "stock2shop-go/api"
	"stock2shop-go/authentication"
	"stock2shop-go/email"
	"stock2shop-go/io"
)

const RUNNING_PORT = "4065"

func main() {

	io.SayHello()

	r := mux.NewRouter()
	myPrefix := "/backend-stock2shop/api"
	//myPrefixWs := "/backend-pmis-ws/api"



	/*==================================================#
	 *========== COMMON SERVICE			    	========#
	 *==================================================#
	 */
	r.Handle(myPrefix+"/common/entity/{module}/{action}", negroni.New(
		negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(api.WsEntityManagement)),
	))
	/*==================================================#
	 *========== AUTHENTICATION SERVICE			 =======#
	 *==================================================#
	 */

	r.Handle(myPrefix+"/email/send", negroni.New(
		negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(email.WsEmail_Send)),
	))

	/*==================================================#
	 *========== AUTHENTICATION SERVICE			 =======#
	 *==================================================#
	 */
	r.HandleFunc(myPrefix+"/user/login", authentication.LoginHandler)
	r.HandleFunc(myPrefix+"/user/new", authentication.WsUser_New)
	r.Handle(myPrefix+"/user/detail", negroni.New(
		negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(authentication.UserDetailHandler)),
	))
	r.Handle(myPrefix+"/user/find", negroni.New(
		negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(authentication.WsUser_Find)),
	))
	r.Handle(myPrefix+"/user/list", negroni.New(
		negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(authentication.WsUser_List)),
	))


	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Origin", "X-Requested-With", "Accept", "X-Token", "x-token", "Content-Type",
		"X-Custom-Header", "UserCode", "OrgCode"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("Listening StaticWebsiteStore...> ", RUNNING_PORT)

	_err := http.ListenAndServe(":"+RUNNING_PORT, handlers.CORS(corsObj, headersOk, methodsOk)(r))

	if _err != nil {
		log.Printf("\x1B[31mServer exit with error: %s\x1B[39m\n", _err)
		//os.Exit(1)
	}

}
