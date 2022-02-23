package handler

import (
	"net/http"
	"surekapi/helper"

	"github.com/gin-gonic/gin"
)

func GetKecepatan(c *gin.Context) {
	kecepatan := []map[string]interface{}{
		{
			"id":   1,
			"name": "Biasa",
		}, {
			"id":   2,
			"name": "Penting",
		}, {
			"id":   3,
			"name": "Segera",
		}, {
			"id":   4,
			"name": "Amat Segera/Kilat",
		},
	}

	response := helper.APIResponse("Master Kecepatan", http.StatusOK, true, kecepatan)
	c.JSON(http.StatusOK, response)
}
