package main

import (
	"net"
	"net/smtp"
	"strings"
)

func parseDomain(email string) string {
	coll := strings.Split(email, "@")
	if len(coll) == 2 {
		return coll[1]
	}
	return ""
}

// VerifyEmail checks an email by using SMTP Callback Verification method
//
// Reference:
// http://en.wikipedia.org/wiki/Callback_verification
func VerifyEmail(email string) (isValid bool, err error) {
	mx, err := net.LookupMX(parseDomain(email))
	if err != nil {
		return
	}
	c, err := smtp.Dial(mx[0].Host + ":25")
	if err != nil {
		return
	}
	err = c.Hello("verify-email.org")
	if err != nil {
		return
	}
	err = c.Mail("check@verify-email.org")
	if err != nil {
		return
	}
	if err := c.Rcpt(email); err != nil {
		return false, nil
	} else {
		return true, nil
	}
}
