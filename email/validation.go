package email

import "fmt"
import "github.com/badoux/checkmail"

func ValidationFormat(address string) (bool, error) {
	err := checkmail.ValidateFormat(address)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
func ValidationDomain(address string) (bool, error) {
	err := checkmail.ValidateHost(address)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
func ValidationUser(address string) (bool, error) {
	err := checkmail.ValidateHost(address)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
		return false, smtpErr
	}
	return true, nil
}
