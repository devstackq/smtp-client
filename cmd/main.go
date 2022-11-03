package main

import (
	"log"

	"github.com/devstackq/smtp-mailer/config"
	handler "github.com/devstackq/smtp-mailer/internal/handler/http"
	"github.com/devstackq/smtp-mailer/internal/repository"
	"github.com/devstackq/smtp-mailer/internal/service"
)

func main() {

	// var wg sync.WaitGroup
	// wg.Add(2)

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

	// go func() {
	// 	que.ClientMailer() //cleint run
	// 	wg.Done()
	// }()

	// go func() {
	// 	que.WorkerMailer(mailer) //wait client
	// 	wg.Done()
	// }()

}
