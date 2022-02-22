package handler

import (
	"net/http"
	"surekapi/helper"
	"surekapi/pejabat"
	"surekapi/user"

	"github.com/gin-gonic/gin"
)

type pejabatHandler struct {
	service pejabat.Service
}

func NewPejabatHandler(service pejabat.Service) *pejabatHandler {
	return &pejabatHandler{service}
}

func (h *pejabatHandler) GetByUnitKerjaID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	kategori := c.Query("kategori")
	masterPejabat, err := h.service.GetByUnitKerjaID(currentUser.UnitKerjaID, kategori)
	if err != nil {
		response := helper.APIResponse("Gagal mengambil data pejabat", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pejabat.FormatMultiplePejabat(masterPejabat)
	response := helper.APIResponse("Master Pejabat", http.StatusOK, true, formatter)
	c.JSON(http.StatusOK, response)
}
