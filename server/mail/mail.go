package mail

import (
	"fmt"
	"os"
)

type emailTemplate string

// ForgotPasswordTemplate is the ID for email template
const ForgotPasswordTemplate = emailTemplate("forgot-password-template")

// ContactTemplate is the ID for contact email template
const ContactTemplate = emailTemplate("contact-template")

// NeedResponseTemplate sis the ID for need response template
const NeedResponseTemplate = emailTemplate("need-response-template")

// Mailer is a implementation to send emails
type Mailer interface {
	SendEmail(EmailParams) error
}

// EmailParams struct with emails infos
type EmailParams struct {
	To        string
	CC        string
	Subject   string
	Template  emailTemplate
	Variables map[string]string
}

// Connect - Create and return a dialer
func Connect() (Mailer, error) {
	method := os.Getenv("MAIL_METHOD")

	switch method {
	case "sendgrid":
		return SendgridConnect()

	case "smtp":
		return SMTPConnect()

	default:
		return nil, fmt.Errorf("mail method %s is not known", method)
	}
}
