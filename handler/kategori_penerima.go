package handler

import (
	"net/http"
	"surekapi/helper"

	"github.com/gin-gonic/gin"
)

func GetKategoriPenerima(c *gin.Context) {
	kategoriPenerima := []map[string]interface{}{
		{
			"id":   1,
			"name": "EXTERNAL",
		}, {
			"id":   2,
			"name": "INTERNAL",
		}, {
			"id":   3,
			"name": "INTERNAL UNIT KERJA",
		},
	}

	response := helper.APIResponse("Master Naskah", http.StatusOK, true, kategoriPenerima)
	c.JSON(http.StatusOK, response)
}
