package email

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"strconv"

	"github.com/wneessen/go-mail"

	"server/config"
)

// Embed the entire templates directory into the binary
//
//go:embed templates/*.tmpl
var tplFS embed.FS

type EmailService struct {
	client *mail.Client
	cfg    *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {

	emailService, err := Init(cfg)
	if err != nil {
		log.Fatalf("failed to create email service: %s", err)
	}

	return emailService
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

	client, err := mail.NewClient(
		cfg.SMTP.Host,
		mail.WithSMTPAuth(auth),
		mail.WithUsername(cfg.SMTP.User),
		mail.WithPassword(cfg.SMTP.Password),
		mail.WithPort(port),
		mail.WithTLSPolicy(tlspolicy),
	)
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	return &EmailService{client: client, cfg: cfg}, nil
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

	log.Println("Sending plain text email to", to)
	return s.client.DialAndSend(msg)
}

func (s *EmailService) SendHTMLTemplate(to, subject, tplName string, data interface{}) error {
	msg := mail.NewMsg()
	msg.Subject(subject)

	if err := msg.From(s.cfg.SMTP.From); err != nil {
		return err
	}
	if err := msg.To(to); err != nil {
		return err
	}

	tpl, err := template.ParseFS(tplFS, "templates/"+tplName, "templates/emailStyles.tmpl")
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer
	if err := tpl.ExecuteTemplate(&bodyBuffer, "main", data); err != nil {
		return err
	}

	msg.SetBodyString(mail.TypeTextHTML, bodyBuffer.String())

	log.Println("Sending HTML email to", to, "using template:", tplName)
	return s.client.DialAndSend(msg)
}
