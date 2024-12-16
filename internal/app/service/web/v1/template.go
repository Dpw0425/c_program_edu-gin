package service

import (
	"c_program_edu-gin/resource"
	"c_program_edu-gin/utils/email"
)

var _ ITemplateService = (*TemplateService)(nil)

type ITemplateService interface {
	LoadTemplate(data map[string]string) ([]byte, error)
}

type TemplateService struct {
}

func (t *TemplateService) LoadTemplate(data map[string]string) ([]byte, error) {
	template, err := resource.Template().ReadFile("templates/email/verify_code.tmpl")
	if err != nil {
		return nil, err
	}

	return email.RenderTemplate(template, data)
}
