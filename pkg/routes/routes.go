package routes

import (
	"github.com/lerryjay/email-service/pkg/config"
)

type Smtp struct {
	Host     string
	Port     int16
	Username string
	Password string
}

func Init(config config.Config) Smtp {
	return Smtp{Host: config.SMTP_HOST, Port: config.SMTP_PORT, Username: config.SMTP_USER, Password: config.SMTP_PWD}
}
