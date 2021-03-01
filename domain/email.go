package domain


type EmailRequest struct {
	Org        string
	From       string
	To         string
	Subject    string
	Body       string
	Files      []string
	ReplyTo    string
	SenderName string
}
