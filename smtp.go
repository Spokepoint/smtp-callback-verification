package main

import (
	"log"
	"net"
	"net/smtp"
	"os"
)

func main() {
	domain := os.Args[1]
	account := os.Args[2]
	mx, _ := net.LookupMX(domain)
	host := mx[0].Host
	log.Println("Found MX Host: " + host)
	c, err := smtp.Dial(host + ":25")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Hello("verify-email.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Mail("check@verify-email.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt(account + "@" + domain); err != nil {
		log.Println("Fail")
	} else {
		log.Println("Success")
	}
}
