package service

import (
	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/devstackq/smtp-mailer/internal/repository"
)

type TemplateService struct {
	filePath string
	repo     repository.TemplateRepository
}

func NewTemplate(filepath string, tmplRepo repository.TemplateRepository) *TemplateService {
	return &TemplateService{
		filePath: filepath,
		repo:     tmplRepo,
	}
}
func (t *TemplateService) GetListTemplate() ([]models.Template, error) {
	return t.repo.GetListTemplates()
}

func (t *TemplateService) GetTempalteByID(id string) (*models.Template, error) {
	return t.repo.GetTemplateById(id)
}

func (t *TemplateService) Create(bodyPage models.Template) error {
	// tmpl := template.New("example.html")
	// t, err := tmpl.Parse(bodyPage)
	if err := t.repo.Create(bodyPage); err != nil {
		return err
	}
	return nil
}
