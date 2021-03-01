package util

import (
	"stock2shop-go/email"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckIfDateIsInRange_InputString(date1, date2, dateCheck string) bool {
	if dateCheck == date1 || dateCheck == date2 {
		return true
	}
	dt1 := StringToTime(date1)
	dt2 := StringToTime(date2)
	dtCheck := StringToTime(dateCheck)
	start := dt1
	end := dt2
	check := dtCheck
	if date1 == strings.Split(time.Now().String(), " ")[0] {
		return true
	}
	if date2 == strings.Split(time.Now().String(), " ")[0] {
		return true
	}
	return check.After(start) && check.Before(end)
}
func StringToTime(date string) time.Time {
	year, month, day := buidMonthDayYear(date)
	start := time.Date(year, time.Month(month), day, 9, 0, 0, 0, time.UTC)
	return start
}
func buidMonthDayYear(dateIn string) (int, int, int) {
	arr := strings.Split(dateIn, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])
	return year, month, day
}
func StringTimeToDateTime(date, inTime string) time.Time {
	year, month, day := buidMonthDayYear(date)
	hour, minute, second := buildHourMinute(inTime)
	start := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	return start
}
func buildHourMinute(inTime string) (int, int, int) {
	arr := strings.Split(inTime, ":")
	hour, _ := strconv.Atoi(arr[0])
	minute, _ := strconv.Atoi(arr[1])
	var second int
	if len(arr) >= 3 {
		second, _ = strconv.Atoi(arr[2])
	}
	return hour, minute, second
}

func SendEmailNotificationPolicy(emailAddress string, subject string, bodyText string, bodyParams map[string]string) {
	if boo, _ := email.ValidationUser(emailAddress); !boo {
		//sorry we can't send email to unresolvable host or wrong email
		fmt.Println("sorry we can't send email to unresolvable host or wrong email")
		return
	}

	var buildThTdHtml = func(key, val string, tag string) string {
		return fmt.Sprintf("<tr><%v>%v</%v><%v>%v</%v></tr>", tag, key, tag, tag, val, tag)
	}
	//todo build your email body
	body := bodyText
	if len(bodyParams) > 1 {
		body = body + fmt.Sprintf("<table border=1>")
		for key, val := range bodyParams {
			body = body + buildThTdHtml(key, val, "td")
		}
		body = body + fmt.Sprintf("</table>")
	}
	body = body + fmt.Sprintf("<p>Please use the Policy Number provide above in any transaction with us</p")
	body = body + fmt.Sprintf("<p>An email with authentication detail will be send to your contact email</p")
	body = body + fmt.Sprintf("<p>For any query please contact  %v by email at: %v  or phone number: %v </p>", ContactName, adminEmail, ContactNumber)
	body = body + fmt.Sprintf("<p>Thank you</p>")

	var myEmail email.MyEmail
	myEmail.FromEmail = "easidoc@easipath.com"
	myEmail.Receiver = emailAddress
	myEmail.Subject = subject
	myEmail.ReplyTo = "no-reply@easipath.com"
	myEmail.Sender = "info@easipath.com"
	myEmail.FromCompany = "pos"
	myEmail.Body = body
	myEmail.Send()
}

func cleanData(inStr string) string {
	inStr = strings.Trim(inStr, " ")
	inStr = strings.ToLower(inStr)
	return inStr
}
