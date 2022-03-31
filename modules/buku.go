package modules

type Buku struct {
	KodeBuku    string `json:"kode_buku" bson:"kode_buku"`
	Title       string `json:"title" bson:"title"`
	Price       int    `json:"price" bson:"price"`
	Description string `json:"description" bson:"description"`
}
