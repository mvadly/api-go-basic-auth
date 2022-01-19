package partner

import (
	"api-klasterku-partner/config"
	"api-klasterku-partner/entity"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var sql = config.ConnectDB().Debug()

func GetProcessedData(id, periode_data string) ([]entity.TableClusterProcessed, error) {
	var data = []entity.TableClusterProcessed{}
	err := sql.Where("id_cluster = ? AND periode_data = ?",
		id,
		periode_data).
		Find(&data).Error
	return data, err

}

func SaveProcessedData(post RequestClusterProcessed) (string, error) {
	idData, err := uuid.NewUUID()
	if err != nil {
		return "error", err
	}
	ID := strings.ReplaceAll(idData.String(), "-", "")
	var kodeUker, _ = strconv.Atoi(post.KodeUker)
	var totalAnggota, _ = strconv.Atoi(post.TotalAnggota)
	var agenBrilink, _ = strconv.Atoi(post.AgenBrilink)
	var trxAgenBrilink, _ = strconv.ParseFloat(post.TrxAgenBrilink, 64)
	var nominalFbi, _ = strconv.ParseFloat(post.NominalFbi, 64)
	var totalRekeningSimpanan, _ = strconv.Atoi(post.TotalRekeningSimpanan)
	var totalSaldoSimpanan, _ = strconv.ParseFloat(post.TotalSaldoSimpanan, 64)
	var totalRekeningPinjaman, _ = strconv.Atoi(post.TotalRekeningPinjaman)
	var totalSaldoPinjaman, _ = strconv.ParseFloat(post.TotalSaldoPinjaman, 64)
	var kolektabilitas, _ = strconv.ParseFloat(post.Kolektibilitas, 32)
	var periodeData, _ = strconv.Atoi(post.PeriodeData)

	var data = entity.TableClusterProcessed{
		KodeUker:              kodeUker,
		IdCluster:             post.IdKlaster,
		IdKategoriUsaha:       post.IdKategoriUsaha,
		NamaKategoriUsaha:     post.KategoriUsaha,
		TotalAnggota:          totalAnggota,
		AgenBrilink:           agenBrilink,
		TrxAgenBrilink:        trxAgenBrilink,
		NominalFbi:            nominalFbi,
		TotalRekeningSimpanan: totalRekeningSimpanan,
		TotalSaldoSimpanan:    totalSaldoSimpanan,
		TotalRekeningPinjaman: totalRekeningPinjaman,
		TotalSaldoPinjaman:    totalSaldoPinjaman,
		Kolektabilitas:        float32(kolektabilitas),
		PeriodeData:           periodeData,
	}
	exist, err := GetProcessedData(post.IdKlaster, post.PeriodeData)
	if len(exist) > 0 {
		data.UpdatedAt = time.Now().Local()
		data.UpdatedBy = post.UsernameBy
		err := sql.Where("id_cluster = ? AND periode_data = ?",
			post.IdKlaster,
			post.PeriodeData).
			Updates(&data).Error
		config.CloseDB()
		return "updated", err
	} else {
		data.IDData = ID
		data.CreatedBy = post.UsernameBy
		err := sql.Create(&data).Error
		config.CloseDB()
		return "created", err
	}

}
