package mail

import (
	"fmt"
	"github.com/mewa/wuff/config"
	"net/smtp"
)

func SendEmail(body string, config *config.Config) error {
	auth := smtp.PlainAuth("", config.Smtp.User, config.Smtp.Password, config.Smtp.Server)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Smtp.Server, config.Smtp.Port),
		auth,
		config.Smtp.Sender,
		[]string{config.Email},
		[]byte(body),
	)

	return err
}
