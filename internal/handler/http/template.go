package handler

import (
	"net/http"

	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTemplate(ctx *gin.Context) {

	var request models.Template

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.tmplSrv.Create(request); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) GetListTemplate(ctx *gin.Context) {

	tempaltes, err := h.tmplSrv.GetListTemplate()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, tempaltes)
}

func (h *Handler) GetTemplateByID(ctx *gin.Context) {
	var request models.Template

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	template, err := h.tmplSrv.GetTempalteByID(request.ID.Hex())
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, template)
}
