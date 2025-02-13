package email

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func SendEmail(to string, subject string, body string) error {
	from := "cookingmaster.cloudx@gmail.com"
	password := "frqe ovcm xoeq xpmf"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := LoginAuth(from, password)
	msg := "To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n"
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		logx.Errorf("smtp.SendMail err: %v", err)
	}
	return err
}
