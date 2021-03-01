package email

import (
	"crypto/tls"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"os"
	"stock2shop-go/constant"
	"strconv"
	"strings"
	"time"
)

type MyEmail struct {
	Username    string
	Sender      string
	Receiver    string
	Subject     string
	Body        string
	ReplyTo     string
	FromEmail   string
	FromCompany string

	HasAttached  string
	Attaches     string
	AttachedType string
}

func (obj *MyEmail) Send() {
	port, _ := strconv.Atoi(SMTP_PORT)
	d := gomail.NewDialer(SMTP_SERVER, port, SMTP_USERNAME, SMTP_PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage(gomail.SetCharset("ISO-8859-1"), gomail.SetEncoding(gomail.Base64))

	fromName := fmt.Sprintf("%s<%s>", obj.FromCompany, obj.FromEmail)

	fmt.Println("fromName :)--> ", fromName)

	m.SetHeaders(map[string][]string{
		"From":    {fromName},
		"To":      {obj.Receiver},
		"Subject": {obj.Subject},
	})
	m.SetBody("text/html", obj.Body)
	/*
		LET MAKE ATTACHMENT REQUE
	*/

	CreateFolderUploadIfNotExist(DIR_TEMP_ATTACHED_DOWNLOAD)

	var lsitdelete []string
	attachedList := strings.Split(obj.Attaches, ";")

	if obj.HasAttached == "yes" {
		if obj.AttachedType == "local" {
			fmt.Println("LOCAL FILE ATTACH >>>> ", obj.AttachedType)
			for _, filename := range attachedList {
				m.Attach(filename)
			}
		}

		fmt.Println("obj.AttachedType>>>> ", obj.AttachedType)

		if obj.AttachedType == "external" {
			for _, fileUrl := range attachedList {
				mybite, filename := global.GetHttpFileContent2(fileUrl)
				attFile := fmt.Sprintf("./%s/%s", constant.DIR_TEMP_ATTACHED_DOWNLOAD, filename)

				ioutil.WriteFile(attFile, mybite, 0644)

				fmt.Println("external  FILE ATTACH >>>> ", attFile)
				m.Attach(attFile)
				timer := time.NewTimer(1 * time.Second)
				lsitdelete = append(lsitdelete, attFile)
				<-timer.C
			}
		}
	}

	err := d.DialAndSend(m)
	fmt.Println("@-=> EMAIL SEND REPORT : ", err)

	timer2 := time.NewTimer(time.Second * 60)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
		for _, fname := range lsitdelete {
			var err = os.Remove(fname)
			global.CheckError(err)
		}
	}()

}

func TemplateEmailGeneral(name, title, msg string) string {
	return `<div style="background-color:#fff;margin:0 auto 0 auto;padding:30px 0 30px 0;color:#4f565d;font-size:13px;line-height:20px;font-family:'Helvetica Neue',Arial,sans-serif;text-align:left">
    <center>
        <table style="width:550px;text-align:center">
            <tbody>
            <tr>
                <td style="padding:0 0 20px 0;border-bottom:1px solid #e9edee">
                    <a href="https://epwp.easipath.com" style="display:block;margin:0 auto" target="_blank"
                       data-saferedirecturl="https://epwp.easipath.com/">
                        <img insurance="https://www.careers24.com/_resx/imageresource/E912F0C5178A352720D82F1EF747CEF5EDC23AC2-366515-400-200-0"
                             alt="MACROBERT logo" style="border:0px" class="CToWUd" width="200" height="100">
                    </a>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0">
                    <p style="color:#1d2227;line-height:28px;font-size:22px;margin:12px 10px 20px 10px;font-weight:400">
                        Hi ` + name + `.
					</p>
					<p style="color:#1d2227;line-height:18px;font-size:17px;margin:12px 10px 20px 10px;font-weight:200">
                         ` + title + `.
					</p>
                    <p style="margin:0 10px 10px 10px;padding:0">
						` + msg + `
					</p>
                </td>
            </tr>
            <tr>
            <tr>
                <td colspan="2" style="padding:30px 0 0 0;border-top:1px solid #e9edee;color:#9b9fa5">
                    If you have any questions you can contact us at <a style="color:#666d74;text-decoration:none"
                                                                       href="mailto:` + constant.MAIN_CONTACT_EMAIL + `" target="_blank">` + constant.MAIN_CONTACT_EMAIL + `</a>
                </td>
            </tr>
            </tr>
            </tbody>
        </table>
    </center>
    <img insurance="https://ci6.googleusercontent.com/proxy/4HHdWNfwIuYk-bX22_jJ-eSs1ihG4pVGzxvkKHUvibCbS1bg4yVBKBS2XTy_i5APCw0mXUz2Wghzlu3-UOOIZ8dg_FdBHGokaZ4CgUX3pR0xztbUh6VSXyjA-9GDaWQocplLpf56OnNfzGZNrym6QupOHQrd_gDAjPQylA=s0-d-e1-ft#https://secure.fastclick.net/w/tre?ad_id=35399;evt=27936;cat1=38343;cat2=38344;rand=636580731486389038"
         class="CToWUd" width="1" border="0" height="\&quot;1&quot;">
    <div class="yj6qo"></div>
    <div class="adL">
    </div>
</div>
`
}

func TemplateEmailVerification(name, tokenRegister string) string {
	return `<div style="background-color:#fff;margin:0 auto 0 auto;padding:30px 0 30px 0;color:#4f565d;font-size:13px;line-height:20px;font-family:'Helvetica Neue',Arial,sans-serif;text-align:left">
    <center>
        <table style="width:550px;text-align:center">
            <tbody>
            <tr>
                <td style="padding:0 0 20px 0;border-bottom:1px solid #e9edee">
                    <a href="` + constant.REDIRCT_ACTIVATION_URL + `" style="display:block;margin:0 auto" target="_blank"
                       data-saferedirecturl="https://epwp.easipath.com/">
                        <img insurance="https://www.careers24.com/_resx/imageresource/E912F0C5178A352720D82F1EF747CEF5EDC23AC2-366515-400-200-0"
                             alt="MACROBERT logo" style="border:0px" class="CToWUd" width="200" height="100">
                    </a>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0">
                    <p style="color:#1d2227;line-height:28px;font-size:22px;margin:12px 10px 20px 10px;font-weight:400">
                        Hi ` + name + `, it's great to meet you.</p>
                    <p style="margin:0 10px 10px 10px;padding:0">We'd like to make sure we got your email address
                        right.</p>
                    <p>
                        <a style="display:inline-block;text-decoration:none;padding:15px 20px;background-color:#2baaed;border:1px solid #2baaed;border-radius:3px;color:#fff;font-weight:bold"
                           href="` + constant.ACTIVATION_USER_SERVER + `/` + tokenRegister + `" target="_blank"
                           data-saferedirecturl="https://www.google.com/url?hl=en&amp;q=https://login.xero.com/c?token%3DxGV-oDDjy4Q-zfWTT7Vqmp&amp;source=gmail&amp;ust=1524216763316000&amp;usg=AFQjCNHyCcYORHHVfcv5drZL5JvUq1JZ0g">
 							click this to verify your email.</a>
                    </p>
                </td>
            </tr>
            <tr>
            <tr>
                <td colspan="2" style="padding:30px 0 0 0;border-top:1px solid #e9edee;color:#9b9fa5">
                    If you have any questions you can contact us at <a style="color:#666d74;text-decoration:none"
                                                                       href="mailto:` + constant.SENDER_EMAIL + `" target="_blank">` + constant.SENDER_EMAIL + `</a>
                </td>
            </tr>
            </tr>
            </tbody>
        </table>
    </center>
    <img insurance="https://ci6.googleusercontent.com/proxy/4HHdWNfwIuYk-bX22_jJ-eSs1ihG4pVGzxvkKHUvibCbS1bg4yVBKBS2XTy_i5APCw0mXUz2Wghzlu3-UOOIZ8dg_FdBHGokaZ4CgUX3pR0xztbUh6VSXyjA-9GDaWQocplLpf56OnNfzGZNrym6QupOHQrd_gDAjPQylA=s0-d-e1-ft#https://secure.fastclick.net/w/tre?ad_id=35399;evt=27936;cat1=38343;cat2=38344;rand=636580731486389038"
         class="CToWUd" width="1" border="0" height="\&quot;1&quot;">
    <div class="yj6qo"></div>
    <div class="adL">
    </div>
</div>
`
}

func TemplateEmailForgotPasswordRequest(name, token string) string {
	return `<div style="background-color:#fff;margin:0 auto 0 auto;padding:30px 0 30px 0;color:#4f565d;font-size:13px;line-height:20px;font-family:'Helvetica Neue',Arial,sans-serif;text-align:left">
    <center>
        <table style="width:550px;text-align:center">
            <tbody>
            <tr>
                <td style="padding:0 0 20px 0;border-bottom:1px solid #e9edee">
                    <a href="https://epwp.easipath.com" style="display:block;margin:0 auto" target="_blank"
                       data-saferedirecturl="https://epwp.easipath.com/">
                        <img insurance="https://www.careers24.com/_resx/imageresource/E912F0C5178A352720D82F1EF747CEF5EDC23AC2-366515-400-200-0"
                             alt="MACROBERT logo" style="border:0px" class="CToWUd" width="200" height="100">
                    </a>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0">
                    <p style="color:#1d2227;line-height:28px;font-size:22px;margin:12px 10px 20px 10px;font-weight:400">
                        Hi  ` + name + ` !</p>

                    <p style="margin:0 10px 10px 10px;padding:0">
                        We received a request to change your password.<br/>
						Please click on the link below to reset your password, if it s not you ignore this email.
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        <a href="` + constant.REDIRCT_ACTIVATION_URL + `/forgot-password-change/` + token + `">Click Here to Reset</a>
                    </p>

                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0 0 0;border-top:1px solid #e9edee;color:#9b9fa5">
                    If you have any questions you can contact us at <a style="color:#666d74;text-decoration:none"
                                                                       href="mailto:` + constant.MAIN_CONTACT_EMAIL + `" target="_blank">` + constant.MAIN_CONTACT_EMAIL + `</a>
                </td>
            </tr>
            </tbody>
        </table>
    </center>
    <img insurance="https://ci6.googleusercontent.com/proxy/4HHdWNfwIuYk-bX22_jJ-eSs1ihG4pVGzxvkKHUvibCbS1bg4yVBKBS2XTy_i5APCw0mXUz2Wghzlu3-UOOIZ8dg_FdBHGokaZ4CgUX3pR0xztbUh6VSXyjA-9GDaWQocplLpf56OnNfzGZNrym6QupOHQrd_gDAjPQylA=s0-d-e1-ft#https://secure.fastclick.net/w/tre?ad_id=35399;evt=27936;cat1=38343;cat2=38344;rand=636580731486389038"
         class="CToWUd" width="1" border="0" height="\&quot;1&quot;">
    <div class="yj6qo"></div>
    <div class="adL">
    </div>
</div>`
}

func ActivationConfirmationEmailBody(name, username, password string) string {

	return `<div style="background-color:#fff;margin:0 auto 0 auto;padding:30px 0 30px 0;color:#4f565d;font-size:13px;line-height:20px;font-family:'Helvetica Neue',Arial,sans-serif;text-align:left">
    <center>
        <table style="width:550px;text-align:center">
            <tbody>
            <tr>
                <td style="padding:0 0 20px 0;border-bottom:1px solid #e9edee">
                    <a href="https://epwp.easipath.com" style="display:block;margin:0 auto" target="_blank"
                       data-saferedirecturl="https://epwp.easipath.com/">
                        <img insurance="https://www.careers24.com/_resx/imageresource/E912F0C5178A352720D82F1EF747CEF5EDC23AC2-366515-400-200-0"
                             alt="MACROBERT logo" style="border:0px" class="CToWUd" width="200" height="100">
                    </a>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0">
                    <p style="color:#1d2227;line-height:28px;font-size:22px;margin:12px 10px 20px 10px;font-weight:400">
                        Hi ` + name + ` !</p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        We'd like to let you know that your account have been activate.
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Below is your login credentials
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Username : <b>` + username + `</b>
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Password : <b>` + password + `</b>
                    </p>
                    <h3 style="margin:0 10px 10px 10px;padding:0">24 hours</h3>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        <a href="` + constant.WEBSITE_HOME + `">Click here to login</a>
                    </p>

                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0 0 0;border-top:1px solid #e9edee;color:#9b9fa5">
                    If you have any questions you can contact us at <a style="color:#666d74;text-decoration:none"
                                                                       href="mailto:support@xero.com" target="_blank">support@pos.com</a>
                </td>
            </tr>
            </tbody>
        </table>
    </center>
    <img insurance="https://ci6.googleusercontent.com/proxy/4HHdWNfwIuYk-bX22_jJ-eSs1ihG4pVGzxvkKHUvibCbS1bg4yVBKBS2XTy_i5APCw0mXUz2Wghzlu3-UOOIZ8dg_FdBHGokaZ4CgUX3pR0xztbUh6VSXyjA-9GDaWQocplLpf56OnNfzGZNrym6QupOHQrd_gDAjPQylA=s0-d-e1-ft#https://secure.fastclick.net/w/tre?ad_id=35399;evt=27936;cat1=38343;cat2=38344;rand=636580731486389038"
         class="CToWUd" width="1" border="0" height="\&quot;1&quot;">
    <div class="yj6qo"></div>
    <div class="adL">
    </div>
</div>`
}

func RegistrationAccountInProgress(name, username, password string) string {

	return `<div style="background-color:#fff;margin:0 auto 0 auto;padding:30px 0 30px 0;color:#4f565d;font-size:13px;line-height:20px;font-family:'Helvetica Neue',Arial,sans-serif;text-align:left">
    <center>
        <table style="width:550px;text-align:center">
            <tbody>
            <tr>
                <td style="padding:0 0 20px 0;border-bottom:1px solid #e9edee">
                    <a href="` + constant.REDIRCT_ACTIVATION_URL + `" style="display:block;margin:0 auto" target="_blank"
                       data-saferedirecturl="https://epwp.easipath.com/">
                        <img insurance="https://www.careers24.com/_resx/imageresource/E912F0C5178A352720D82F1EF747CEF5EDC23AC2-366515-400-200-0"
                             alt="MACROBERT logo" style="border:0px" class="CToWUd" width="200" height="100">
                    </a>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0">
                    <p style="color:#1d2227;line-height:28px;font-size:22px;margin:12px 10px 20px 10px;font-weight:400">
                        Hi ` + name + ` !</p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        We'd like to let you know that your account have been created.
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Please wait for Admin decision, an email will be send to you when approve or reject
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Below is your login credentials
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Username : <b>` + username + `</b>
                    </p>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        Password : <b>*********</b>
                    </p>
                    <h3 style="margin:0 10px 10px 10px;padding:0">24 hours</h3>
                    <p style="margin:0 10px 10px 10px;padding:0">
                        <a href="` + constant.WEBSITE_HOME + `">Click here to login</a>
                    </p>

                </td>
            </tr>
            <tr>
                <td colspan="2" style="padding:30px 0 0 0;border-top:1px solid #e9edee;color:#9b9fa5">
                    If you have any questions you can contact us at <a style="color:#666d74;text-decoration:none"
                                                                       href="mailto:support@xero.com" target="_blank">support@pmis.com</a>
                </td>
            </tr>
            </tbody>
        </table>
    </center>
    <img insurance="https://ci6.googleusercontent.com/proxy/4HHdWNfwIuYk-bX22_jJ-eSs1ihG4pVGzxvkKHUvibCbS1bg4yVBKBS2XTy_i5APCw0mXUz2Wghzlu3-UOOIZ8dg_FdBHGokaZ4CgUX3pR0xztbUh6VSXyjA-9GDaWQocplLpf56OnNfzGZNrym6QupOHQrd_gDAjPQylA=s0-d-e1-ft#https://secure.fastclick.net/w/tre?ad_id=35399;evt=27936;cat1=38343;cat2=38344;rand=636580731486389038"
         class="CToWUd" width="1" border="0" height="\&quot;1&quot;">
    <div class="yj6qo"></div>
    <div class="adL">
    </div>
</div>`
}
