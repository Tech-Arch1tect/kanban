package email

import (
	"log"
	"server/config"
	"strconv"

	"github.com/wneessen/go-mail"
)

type EmailService struct {
	client *mail.Client
	cfg    *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func Init(cfg *config.Config) (*EmailService, error) {
	var auth mail.SMTPAuthType
	switch cfg.SMTP.Auth {
	case "plain":
		auth = mail.SMTPAuthPlain
	case "login":
		auth = mail.SMTPAuthLogin
	default:
		auth = mail.SMTPAuthNoAuth
	}
	port, err := strconv.Atoi(cfg.SMTP.Port)
	if err != nil {
		log.Fatalf("failed to convert smtp port to int: %s", err)
	}
	tlspolicy := mail.TLSMandatory
	if cfg.SMTP.NoTLS {
		tlspolicy = mail.NoTLS
	}
	client, err := mail.NewClient(cfg.SMTP.Host, mail.WithSMTPAuth(auth),
		mail.WithUsername(cfg.SMTP.User), mail.WithPassword(cfg.SMTP.Password), mail.WithPort(port), mail.WithTLSPolicy(tlspolicy))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	return &EmailService{client: client}, nil
}

func (s *EmailService) SendPlainText(to, subject, body string) error {
	msg := mail.NewMsg()
	msg.Subject(subject)
	if err := msg.From(s.cfg.SMTP.From); err != nil {
		return err
	}
	if err := msg.To(to); err != nil {
		return err
	}
	msg.SetBodyString(mail.TypeTextPlain, body)
	log.Println("Sending email to", to)
	return s.client.DialAndSend(msg)
}
