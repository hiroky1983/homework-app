package repository

import "homework/config"

type Mail interface {
	SendMail(email, tokne string, conf config.Config) error
}
