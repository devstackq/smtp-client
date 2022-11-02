package service

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/devstackq/smtp-mailer/config"
	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/devstackq/smtp-mailer/internal/repository"
)

type Smtp struct {
	filePath       string
	name, password string
	host           string
	port           string
	smtp           smtp.Client
	repoUsr        repository.UserRepository
	repoTmpl       repository.TemplateRepository
}

func NewSmtp(cfg *config.Config, repoUsr repository.UserRepository, repoTmpl repository.TemplateRepository) (Mailer, error) {

	if cfg == nil || repoTmpl == nil || repoUsr == nil {
		return nil, errors.New("empty cofig or repos")
	}
	s := &Smtp{
		filePath: cfg.HtmlFilePath,
		name:     cfg.Username,
		password: cfg.Password,
		host:     cfg.Host,
		port:     cfg.Port,
		repoUsr:  repoUsr,
		repoTmpl: repoTmpl,
	}
	return s, nil

}

func (s *Smtp) auth() smtp.Auth {
	return smtp.PlainAuth("", s.name, s.password, s.host) //remove creds
}

func (s *Smtp) Send(mail models.Mail) error {

	templateDB, err := s.repoTmpl.GetTemplateById(mail.TmplTypeID)
	if err != nil {
		return err
	}

	// users, err := s.repoUsr.GetListUser()
	// if err != nil {
	// 	return err
	// }

	// mock data
	recipients := []string{"upzq@mail.ru", "8akebaev@gmail.com"}

	// for _, user := range users {
	// 	recipients = append(recipients, user.Email)
	// }

	mail.To = recipients

	tmpl := template.New("example.html")
	//set struct to new html
	tmpl, err = tmpl.Parse(templateDB.BodyPage)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	//set date from mail to html
	if err = tmpl.Execute(buf, mail); err != nil {
		return err
	}
	mail.Body = buf.Bytes()

	address := fmt.Sprint(s.host, ":", s.port)
	//send msg to mail
	fmt.Println(string(mail.Body))

	if err = smtp.SendMail(address, s.auth(), mail.From, mail.To, mail.Body); err != nil {
		return err
	}

	return nil
}

/*func newSmptCLient(){
 var conf = &tls.Config{ServerName: host}
 var conn, err = tls.Dial("tcp", addr, conf)
 var cl, err = smtp.NewClient(conn, host)
 err = cl.Auth(auth)
 err = cl.Mail(username)
 err = cl.Rcpt(username)
 var w, err = cl.Data()
 _, err = w.Write(msg)
 err = w.Close()
 err = cl.Quit()
}
*/
