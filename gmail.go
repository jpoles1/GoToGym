package main

import (
	"bytes"
	"html/template"

	gomail "gopkg.in/gomail.v2"
)

func sendEmail(toEmail string, subject string, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", envSMTPSender)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(envSMTPURI, envSMTPPort, envSMTPSender, envSMTPPass)

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

//func sendConfirmationCode(user *User) {
func sendRegistrationEmail(userData *UserDocument) error {
	verificationLink := envBindURL + "/api/verifyemail/" + userData.APIKey

	tmpl := template.New("newuser")
	tmpl, _ = template.ParseFiles("templates/emailregistration.gohtml")
	var tplString bytes.Buffer
	err := tmpl.Execute(&tplString, map[string]string{
		"FirstName":        userData.FirstName,
		"APIKey":           userData.APIKey,
		"VerificationLink": verificationLink,
	})
	errCheck("Rendering email template", err)
	return sendEmail(userData.Email, "GoToGym - Registration", tplString.String())
}
func sendPasswordResetEmail(userData *UserDocument, requestIP string) error {
	verificationLink := envBindURL + "/api/verifyemail/" + userData.APIKey

	tmpl := template.New("newuser")
	tmpl, _ = template.ParseFiles("templates/emailregistration.gohtml")
	var tplString bytes.Buffer
	err := tmpl.Execute(&tplString, map[string]string{
		"FirstName":        userData.FirstName,
		"APIKey":           userData.APIKey,
		"VerificationLink": verificationLink,
	})
	errCheck("Rendering email template", err)
	return sendEmail(userData.Email, "GoToGym - Registration", tplString.String())
}
func sendGymVisitCheckinEmail(visitData GymVisitDocument, userData *UserDocument) error {
	verificationLink := envBindURL + "/api/verifyvisit/" + visitData.ID.Hex() + "/" + userData.APIKey

	tmpl := template.New("visitcheckin")
	tmpl, _ = template.ParseFiles("templates/visitcheckin.gohtml")
	var tplString bytes.Buffer
	err := tmpl.Execute(&tplString, map[string]string{
		"FirstName":        userData.FirstName,
		"StartTime":        visitData.StartTime,
		"EndTime":          visitData.EndTime,
		"VerificationLink": verificationLink,
	})
	errCheck("Rendering email template", err)
	return sendEmail(userData.Email, "GoToGym - Gym Check-In - "+visitData.EndTime, tplString.String())
}
