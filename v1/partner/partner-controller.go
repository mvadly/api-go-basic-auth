package partner

import (
	"api-klasterku-partner/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetDataKlaster(c *gin.Context) {
	c.JSON(200, gin.H{"status": true, "message": "ok"})
}

func ProcessedDataKlaster(c *gin.Context) {
	validate = validator.New()
	validate.RegisterValidation("correct-format", util.ValidateCorrectInput)
	username, _, _ := util.GetUserBasicAuth(c)
	var req = RequestClusterProcessed{
		IdKlaster:             c.PostForm("id_klaster"),
		KodeUker:              c.PostForm("kode_uker"),
		IdKategoriUsaha:       c.PostForm("id_kategori_usaha"),
		KategoriUsaha:         c.PostForm("kategori_usaha"),
		TotalAnggota:          c.PostForm("total_anggota"),
		AgenBrilink:           c.PostForm("agen_brilink"),
		TrxAgenBrilink:        c.PostForm("trx_agen_brilink"),
		NominalFbi:            c.PostForm("nominal_fbi"),
		TotalRekeningSimpanan: c.PostForm("total_rekening_simpanan"),
		TotalSaldoSimpanan:    c.PostForm("total_saldo_simpanan"),
		TotalRekeningPinjaman: c.PostForm("total_rekening_pinjaman"),
		TotalSaldoPinjaman:    c.PostForm("total_saldo_pinjaman"),
		Kolektibilitas:        c.PostForm("kolektibilitas"),
		PeriodeData:           c.PostForm("periode_data"),
		UsernameBy:            username,
	}

	err := validate.Struct(req)
	if err != nil {
		myError := util.GetError(err)
		c.JSON(400, gin.H{"status": false, "message": "validation errors", "errors": myError})

		return
	}

	save, err := SaveProcessedData(req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"status": false, "message": "failed to " + save})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "data klaster " + save})
}

func UploadFile(c *gin.Context) {
	upload, err := util.UploadFile(c, "file", "images/")
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(400, gin.H{"status": false, "message": upload})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": upload})
}
