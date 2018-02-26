package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func sendEmail(toEmail string, subject string, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", envSMTPSender)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(envSMTPURI, envSMPTPPort, envSMTPSender, envSMTPPass)

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

//func sendConfirmationCode(user *User) {
func sendGymVisitCheckin(user *User) error {
	var domainName = envBindURL
	activationLink := user.EmailUUID

	link := domainName + "/emailconfirm/" + activationLink
	message := fmt.Sprintf("Hello %s! <br><br> Thanks for signing up for a Transit Sign Account! <br><br> To activate your account please click the link below: <br><br> <a href=\"%s\">%s</a> <br> <i>If the link does not open automatically please copy and paste it into your browser of choice</i> <br> <br> Thanks! <br> -The Transit Sign Team", user.FirstName, link, link)
	return sendEmail(user.Email, "Transit Server Confirmation!", message)
}
