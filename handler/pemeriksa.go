package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"surekapi/helper"
	"surekapi/user"

	"github.com/gin-gonic/gin"
)

var s []map[string]string

func GetPemeriksa(c *gin.Context) {
	s = []map[string]string{}
	currentUser := c.MustGet("currentUser").(user.User)
	nipPenandatangan := c.Query("nip_penandatangan")
	strukturPemeriksa(currentUser.Nip, nipPenandatangan)
	response := helper.APIResponse("Daftar Pemeriksa", http.StatusOK, true, s)
	c.JSON(http.StatusOK, response)
}

func strukturPemeriksa(nipPembuat string, nipPemeriksa string) (string, bool) {

	nipAwal := nipPembuat
	nipAkhir := nipPemeriksa

	url := "http://simpeg.bkd.sumbarprov.go.id/webapi/pegawai/asn/pimpinan/token/XBnKaywRCrj05m-XXX-v6DXuZ3FFkUgiw45/nip/" + nipAwal
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var resultBody interface{}

	_ = json.Unmarshal([]byte(body), &resultBody)

	responseSimpeg := resultBody.(map[string]interface{})
	result := responseSimpeg["result"].(map[string]interface{})
	unitKerjaID := result["opd_id"].(string)
	nipSimpeg := result["nip"].(string)
	namaAsnSimpeg := result["nama_pns"].(string)
	jabatanIDSimpeg := result["jabatan_id"].(string)
	jabatanAsnSimpeg := result["jabatan_nm"].(string)

	if nipAwal != nipSimpeg {
		nipAwal = nipSimpeg
		data := map[string]string{
			"unit_kerja_id": unitKerjaID,
			"jabatan_id":    jabatanIDSimpeg,
			"jabatan":       jabatanAsnSimpeg,
			"nip":           nipSimpeg,
			"nama":          namaAsnSimpeg,
		}

		s = append(s, data)

		if nipAwal == nipAkhir {
			return "error", false
		}

		strukturPemeriksa(nipAwal, nipAkhir)
	}

	return result["nip"].(string), true
}
