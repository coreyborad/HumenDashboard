package controllers

import (
	"net/http"
	"stock/models"
	"stock/services"

	"github.com/gin-gonic/gin"
)

type XlsxController struct {
	serv *services.XlsxService
}

// NewXlsxController NewXlsxController
func NewXlsxController(service *services.XlsxService) *XlsxController {
	return &XlsxController{
		serv: service,
	}
}

func (c *XlsxController) XlsxAppendRecord(ctx *gin.Context) {
	dataForm := models.XlsxForm{}
	if err := ctx.ShouldBindJSON(&dataForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := c.serv.AppendRecord(dataForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, nil)
}
