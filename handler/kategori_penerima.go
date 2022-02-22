package handler

import (
	"net/http"
	"surekapi/helper"

	"github.com/gin-gonic/gin"
)

func GetKategoriPenerima(c *gin.Context) {
	kategoriPenerima := map[int]string{
		1: "EXTERNAL",
		2: "INTERNAL",
		3: "INTERNAL UNIT KERJA",
	}

	response := helper.APIResponse("Master Naskah", http.StatusOK, true, kategoriPenerima)
	c.JSON(http.StatusOK, response)
}
