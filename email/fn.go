package email

import (
	"os"
)

func CreateFolderUploadIfNotExist(path string) {
	_ = os.MkdirAll(path, os.ModePerm)
}
func SendEmailNotification(in EmailRequest) {
	ServiceSendSampleEmail(in)
}
