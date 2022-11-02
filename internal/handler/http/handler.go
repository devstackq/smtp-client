package handler

import (
	"net/http"

	"github.com/devstackq/smtp-mailer/internal/service"
	"github.com/gin-gonic/gin"
)

// call smtp || queu
type Handler struct {
	usrSrv  *service.UserService
	tmplSrv *service.TemplateService
	mailSrv service.Mailer
}

func New(userSrv *service.UserService, tmplSrv *service.TemplateService, mailer service.Mailer) *Handler {
	return &Handler{
		usrSrv:  userSrv,
		tmplSrv: tmplSrv,
		mailSrv: mailer,
	}
}

func (h *Handler) Register() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/:id", h.GetUserByID)
		user.GET("", h.GetListUser)
		user.POST("", h.CreateUser)
	}

	tmpl := r.Group("/template")
	{
		tmpl.GET("/:id", h.GetTemplateByID)
		tmpl.GET("", h.GetListTemplate)
		tmpl.POST("", h.CreateTemplate)
	}

	mailer := r.Group("/send-mail")
	{
		mailer.POST("", h.SendMail)
	}

	http.ListenAndServe(":8080", r)

}
