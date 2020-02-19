package email

import (
	"bytes"
	"context"
	"text/template"

	"github.com/JesusIslam/goinblue"
)

type SendBlueDataInterface interface {
	SendMail(context.Context) error
	TemplateRender(context.Context, interface{}, string) string
}

var basePath = "./email/templates/"

type SendBlueData struct {
	ToName    string
	ToEmail   string
	Subject   string
	FromEmail string
	Content   string
}

//SendInBlue struct contains secret configuration
type SendInBlue struct {
	EmailSecret string
}

//NewSendInBlue initilize struct dependancy
func NewSendInBlue(emailSecret string) *SendInBlue {
	return &SendInBlue{EmailSecret: emailSecret}
}

//sconst EmailSecret = "RcZ3tf5EDO9b0AFK"

//SendMail send email by using NewSendInBlue
func (secret *SendInBlue) SendMail(ctx context.Context, s SendBlueData) (*goinblue.Response, error) {

	email := &goinblue.Email{
		To: map[string]string{
			s.ToEmail: s.ToName,
		},
		Subject: s.Subject,
		From: []string{
			s.FromEmail, "From",
		},
		Html: s.Content,
	}

	client := goinblue.NewClient(secret.EmailSecret)

	res, err := client.SendEmail(email)

	if err != nil {
		return nil, err
	}
	return res, nil
}

//TemplateRender **
func (s *SendBlueData) TemplateRender(ctx context.Context, data interface{}, TemplateName string) string {

	t1, err := template.ParseFiles(basePath + TemplateName)
	if err != nil {
		return err.Error()
	}

	var tpl bytes.Buffer
	if err := t1.Execute(&tpl, data); err != nil {
		return err.Error()
	}

	result := tpl.String()
	return result
}
