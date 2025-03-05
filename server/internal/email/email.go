package email

import (
	"bytes"
	"embed"
	"html/template"
	"strconv"

	"github.com/wneessen/go-mail"
	"go.uber.org/zap"

	"server/config"
)

// Embed the entire templates directory into the binary
//
//go:embed templates/*.tmpl
var tplFS embed.FS

type EmailService struct {
	client *mail.Client
	cfg    *config.Config
	logger *zap.Logger
}

func NewEmailService(cfg *config.Config, logger *zap.Logger) *EmailService {

	emailService, err := Init(cfg, logger)
	if err != nil {
		logger.Error("failed to create email service", zap.Error(err))
	}

	return emailService
}

func Init(cfg *config.Config, logger *zap.Logger) (*EmailService, error) {
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
		logger.Error("failed to convert smtp port to int", zap.Error(err))
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
		logger.Error("failed to create mail client", zap.Error(err))
	}

	return &EmailService{client: client, cfg: cfg, logger: logger}, nil
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

	s.logger.Info("Sending plain text email to", zap.String("to", to))
	if s.cfg.Environment != "testing" {
		return s.client.DialAndSend(msg)
	}
	return nil
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

	s.logger.Info("Sending HTML email to", zap.String("to", to), zap.String("template", tplName))
	if s.cfg.Environment != "testing" {
		return s.client.DialAndSend(msg)
	}
	return nil
}
