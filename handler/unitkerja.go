package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/unitkerja"
	"surekapi/user"

	"github.com/gin-gonic/gin"
)

type unitKerjaHandler struct {
	service unitkerja.Service
}

func NewUnitKerjaHandler(service unitkerja.Service) *unitKerjaHandler {
	return &unitKerjaHandler{service}
}

func (h *unitKerjaHandler) GetAll(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	masterUnitKerja, err := h.service.GetAll(currentUser.UnitKerjaID)
	if err != nil {
		response := helper.APIResponse("Gagal mengambil data unit kerja", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := unitkerja.FormatMultipleUnitKerja(masterUnitKerja)
	response := helper.APIResponse("Master Unit Kerja", http.StatusOK, true, formatter)
	c.JSON(http.StatusOK, response)
}
