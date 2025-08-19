package transaksi_uang_masuk_ppdb

type TableDataList struct {
	TotalDataPendaftar int   `json:"totalDataPendaftar"`
	JmlDataSudahImport int   `json:"jmlDataSudahImport"`
	JmlDataBelumImport int   `json:"jmlDataBelumImport"`
	Total              int64 `json:"total"`
	Page               int   `json:"page"`
	Last_page          int   `json:"last_page"`
}
