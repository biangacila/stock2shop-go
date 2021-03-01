package email

import (
	"encoding/json"
	"github.com/biangacila/luvungula-go/global"
	"net/http"
)

func WsEmail_Send(w http.ResponseWriter, r *http.Request) {
	_, dataString := global.GetPostedDataMapAndString(r)
	o := EmailRequest{}
	json.Unmarshal([]byte(dataString), &o)
	go SendEmailNotification(o)

	response := make(map[string]interface{})
	response["status"] = "email sent"

	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = o
	my["RESULT"] = response
	global.PublishToReact(w, r, my, 200)
}
