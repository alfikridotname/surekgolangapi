package handler

import (
	"net/http"
	"surekapi/helper"

	"github.com/gin-gonic/gin"
)

func GetKeamanan(c *gin.Context) {
	keamanan := []map[string]interface{}{
		{
			"id":   1,
			"name": "Surat Biasa",
		}, {
			"id":   2,
			"name": "Surat Penting",
		}, {
			"id":   3,
			"name": "Surat Rahasia",
		},
	}

	response := helper.APIResponse("Master Keamanan", http.StatusOK, true, keamanan)
	c.JSON(http.StatusOK, response)
}
