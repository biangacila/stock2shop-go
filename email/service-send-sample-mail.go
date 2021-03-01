package email

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ServiceSendSampleEmail(in EmailRequest) {
	hub := MyEmail{}
	hub.FromEmail = "easidoc@easipath.com"
	if in.From != "" {
		hub.FromEmail = in.From
	}
	hub.Receiver = in.To
	hub.Subject = in.Subject
	hub.Body = in.Body
	hub.ReplyTo = "easidoc@easipath.com"
	if in.ReplyTo != "" {
		hub.ReplyTo = in.ReplyTo
	}
	hub.Sender = in.SenderName
	hub.FromCompany = strings.ToUpper(in.SenderName)

	str, _ := json.Marshal(hub)
	fmt.Println(string(str))
	hub.Send()
}
