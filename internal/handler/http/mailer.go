package handler

import (
	"fmt"
	"net/http"

	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendMail(ctx *gin.Context) {

	var request models.Mail

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.mailSrv.Send(request); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)

}
