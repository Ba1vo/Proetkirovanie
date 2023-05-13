package email

import (
	"errors"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func Email(reciever string, code string) error {
	sender := "2002egor@gmail.com"
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", reciever)
	msg.SetHeader("Subject", "Test msg")
	msg.SetBody("text/html", fmt.Sprintf("<b> Your verification link: localhost:3000/verify?email=%s&c=%s </b>", reciever, code))
	n := gomail.NewDialer("smtp.gmail.com", 587, sender, "wubbkqpidfzzqzvm")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return errors.New("email error")
	}
	return nil
}
