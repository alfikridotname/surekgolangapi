package handler

import (
	"fmt"
	"net/http"
	"surekapi/helper"
	"surekapi/surat"
	"surekapi/user"

	"github.com/gin-gonic/gin"
)

type suratHandler struct {
	service surat.Service
}

func NewSuratHandler(service surat.Service) *suratHandler {
	return &suratHandler{service}
}

func (h *suratHandler) CreateSurat(c *gin.Context) {
	var input surat.BuatSuratInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		fmt.Println(err)
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create surat failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.CreatedBy = currentUser.ID
	input.UpdatedBy = currentUser.ID
	input.CreatedNIP = currentUser.Nip
	input.JabatanID = currentUser.JabatanID
	input.UnitKerjaID = currentUser.UnitKerjaID

	_, err = h.service.CreateSurat(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create surat failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Successfuly Create Surat", http.StatusOK, true, nil)

	c.JSON(http.StatusOK, response)
}
