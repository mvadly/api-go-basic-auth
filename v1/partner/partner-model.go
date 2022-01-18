package partner

type RequestClusterProcessed struct {
	KodeUker              string `form:"kode_uker" validate:"required,numeric,max=10"`
	IdKlaster             string `form:"id_klaster" validate:"required,alphanum,max=64"`
	IdKategoriUsaha       string `form:"id_kategori_usaha" validate:"required,alphanum,max=64"`
	KategoriUsaha         string `form:"kategori_usaha" validate:"required,correct-format,max=100"`
	TotalAnggota          string `form:"total_anggota" validate:"required,numeric"`
	AgenBrilink           string `form:"agen_brilink" validate:"required,numeric,max=10"`
	TrxAgenBrilink        string `form:"trx_agen_brilink" validate:"required,numeric"`
	NominalFbi            string `form:"nominal_fbi" validate:"required,numeric"`
	TotalRekeningSimpanan string `form:"total_rekening_simpanan" validate:"required,numeric"`
	TotalSaldoSimpanan    string `form:"total_saldo_simpanan" validate:"required,numeric"`
	TotalRekeningPinjaman string `form:"total_rekening_pinjaman" validate:"required,numeric"`
	TotalSaldoPinjaman    string `form:"total_saldo_pinjaman" validate:"required,numeric"`
	Kolektibilitas        string `form:"kolektibilitas" validate:"required,numeric,max=5"`
	PeriodeData           string `form:"periode_data" validate:"required,numeric,max=6"`
	UsernameBy            string
}
