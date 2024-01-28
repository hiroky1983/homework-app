package persistence

import (
	"fmt"
	"homework/config"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
}

func NewMail() *Mail {
	return &Mail{}
}

func (m *Mail) SendMail(email, token string, cnf config.Config) error {
	subject := "アカウント本登録のお願い"
	body := "http://localhost:3000/top"
	from := "noreply@example.net"
	receiver := []string{"hirockysan1983@gmail.com"}

	smtpServer := fmt.Sprintf("%s:%d", cnf.SMTPHost, 1025)
	auth := smtp.CRAMMD5Auth(cnf.SMTPUsername, cnf.SMTPPassword)
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(receiver, ","), subject, body))

	if err := smtp.SendMail(smtpServer, auth, from, receiver, msg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}
