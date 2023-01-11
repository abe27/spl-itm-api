package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	"github.com/abe/erp.api/configs"
)

func SendMail(to, subject, body string) (int, string, time.Time) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = configs.MAIL_USERNAME
	headers["Name"] = "ERP Support System."
	headers["To"] = to
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	serverName := fmt.Sprintf("%s:%s", configs.MAIL_SMTP, configs.MAIL_SMTP_PORT)

	auth := smtp.PlainAuth("", configs.MAIL_USERNAME, configs.MAIL_PASSWORD, configs.MAIL_SMTP)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         configs.MAIL_SMTP,
	}

	c, err := smtp.Dial(serverName)
	if err != nil {
		return 500, err.Error(), time.Now()
	}

	c.StartTLS(tlsConfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		return 500, err.Error(), time.Now()
	}

	// To && From
	if err = c.Mail(configs.MAIL_USERNAME); err != nil {
		return 500, err.Error(), time.Now()
	}

	if err = c.Rcpt(to); err != nil {
		return 500, err.Error(), time.Now()
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return 500, err.Error(), time.Now()
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return 500, err.Error(), time.Now()
	}

	err = w.Close()
	if err != nil {
		return 500, err.Error(), time.Now()
	}

	c.Quit()
	return 200, fmt.Sprintf("Send mail to %s is success.", to), time.Now()
}
