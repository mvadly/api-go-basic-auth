package entity

import "time"

type TableClusterProcessed struct {
	IDData                string    `gorm:"column:id_data;type:varchar(100) not null;unique;primaryKey;"`
	KodeUker              int       `gorm:"column:kode_uker;type:int(10) not null" json:"kode_uker"`
	IdCluster             string    `gorm:"column:id_cluster;type:varchar(100) not null" json:"id_klaster"`
	IdKategoriUsaha       string    `gorm:"column:id_kategori_usaha;type:varchar(50) not null" json:"id_kategori_usaha"`
	NamaKategoriUsaha     string    `gorm:"column:nama_kategori_usaha;type:text not null" json:"kategori_usaha"`
	TotalAnggota          int       `gorm:"column:total_anggota;type:int(10) not null" json:"total_anggota"`
	AgenBrilink           int       `gorm:"column:agen_brilink;type:int(10) not null" json:"agen_brilink"`
	TrxAgenBrilink        float64   `gorm:"column:trx_agen_brilink;type:double not null" json:"trx_agen_brilink"`
	NominalFbi            float64   `gorm:"column:nominal_fbi;type:double not null" json:"nominal_fbi"`
	TotalRekeningSimpanan int       `gorm:"column:total_rekening_simpanan;type:int(10) not null" json:"total_rekening_simpanan"`
	TotalSaldoSimpanan    float64   `gorm:"column:total_saldo_simpanan;type:double not null" json:"total_saldo_simpanan"`
	TotalRekeningPinjaman int       `gorm:"column:total_rekening_pinjaman;type:int(10) not null" json:"total_rekening_pinjaman"`
	TotalSaldoPinjaman    float64   `gorm:"column:total_saldo_pinjaman;type:double not null" json:"total_saldo_pinjaman"`
	Kolektabilitas        float32   `gorm:"column:kolektabilitas;type:double not null" json:"kolektabilitas"`
	PeriodeData           int       `gorm:"column:periode_data;type:int(10) not null" json:"periode_data"`
	CreatedAt             time.Time `gorm:"column:created_at;type:datetime DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy             string    `gorm:"column:created_by;type:varchar(50) not null" json:"created_by"`
	UpdatedAt             time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	UpdatedBy             string    `gorm:"column:updated_by;type:varchar(50) not null" json:"updated_by"`
}

func (TableClusterProcessed) TableName() string {
	return "cluster_processeddata"
}
