package email

import (
	"log"
	"server/config"
	"strconv"

	"github.com/wneessen/go-mail"
)

var c *mail.Client

func Init() error {
	var auth mail.SMTPAuthType
	switch config.CFG.SMTP.Auth {
	case "plain":
		auth = mail.SMTPAuthPlain
	case "login":
		auth = mail.SMTPAuthLogin
	default:
		auth = mail.SMTPAuthNoAuth
	}
	port, err := strconv.Atoi(config.CFG.SMTP.Port)
	if err != nil {
		log.Fatalf("failed to convert smtp port to int: %s", err)
	}
	tlspolicy := mail.TLSMandatory
	if config.CFG.SMTP.NoTLS {
		tlspolicy = mail.NoTLS
	}
	client, err := mail.NewClient(config.CFG.SMTP.Host, mail.WithSMTPAuth(auth),
		mail.WithUsername(config.CFG.SMTP.User), mail.WithPassword(config.CFG.SMTP.Password), mail.WithPort(port), mail.WithTLSPolicy(tlspolicy))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	c = client
	return nil
}

func SendPlainText(to, subject, body string) error {
	msg := mail.NewMsg()
	msg.Subject(subject)
	msg.From(config.CFG.SMTP.From)
	msg.To(to)
	msg.SetBodyString(mail.TypeTextPlain, body)
	log.Println("Sending email to", to)
	return c.DialAndSend(msg)
}
