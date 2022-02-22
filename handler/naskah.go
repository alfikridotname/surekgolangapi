package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/naskah"

	"github.com/gin-gonic/gin"
)

type masterNaskahHandler struct {
	service naskah.Service
}

func NewMasterNaskahHandler(service naskah.Service) *masterNaskahHandler {
	return &masterNaskahHandler{service}
}

func (h *masterNaskahHandler) GetAll(c *gin.Context) {
	masterNaskah, err := h.service.GetAll()
	if err != nil {
		response := helper.APIResponse("Gagal mengambil data naskah", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := naskah.FormatMultipleNaskah(masterNaskah)
	response := helper.APIResponse("Master Naskah", http.StatusOK, true, formatter)
	c.JSON(http.StatusOK, response)
}
