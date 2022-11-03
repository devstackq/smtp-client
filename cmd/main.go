package main

import (
	"log"
	"sync"

	"github.com/devstackq/smtp-mailer/config"
	handler "github.com/devstackq/smtp-mailer/internal/handler/http"
	"github.com/devstackq/smtp-mailer/internal/handler/que"
	"github.com/devstackq/smtp-mailer/internal/repository"
	"github.com/devstackq/smtp-mailer/internal/service"
)

func main() {
	cfg := config.New()
	cfg.Load()

	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repoTmpl := repository.NewTemplateMongo(db)
	srvTmpl := service.NewTemplate(cfg.HtmlFilePath, repoTmpl)

	repoUser := repository.NewUserMongo(db)
	srvUser := service.NewUser(repoUser)

	mailer := service.NewMailer("smtp", cfg, repoUser, repoTmpl)
	h := handler.New(srvUser, srvTmpl, mailer)

	h.Register()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		que.CelerySendMail(mailer)
		wg.Done()
	}()
	wg.Wait()

}
