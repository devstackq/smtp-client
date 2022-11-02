package service

import (
	"github.com/devstackq/smtp-mailer/config"
	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/devstackq/smtp-mailer/internal/repository"
)

type Mailer interface {
	Send(models.Mail) error
}

func NewMailer(mailType string, cfg *config.Config, repoUser repository.UserRepository, repoTmpl repository.TemplateRepository) Mailer {
	if mailType == "smtp" {
		smtp, err := NewSmtp(cfg, repoUser, repoTmpl)
		if err != nil {
			return nil
		}
		return smtp
	}
	return nil
}
