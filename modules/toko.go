package modules

type Toko struct {
	KodeToko string `json:"kode_toko" bson:"kode_toko"`
	NamaToko string `json:"nama_toko" bson:"nama_toko"`
	Domisili string `json:"domisili" bson:"domisili"`
	Buku     []Buku `json:"buku" bson:"buku"`
}
