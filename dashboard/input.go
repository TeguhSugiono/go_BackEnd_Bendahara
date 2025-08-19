package dashboard

type Data_Transaksi struct {
	DetailData interface{}
}

type Total_Detail_Transaksi struct {
	Keterangan string
	Total      float64
}

type returnpostuang struct {
	Kd_group  string
	Nm_group  string
	Sisa_uang float64
}
