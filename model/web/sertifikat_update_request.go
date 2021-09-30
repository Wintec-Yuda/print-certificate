package web

type SertifikatUpdateRequest struct {
	Id   int    `validate:"required"`
	Nama string `validate:"max=200,min=1" json:"nama"`
	Penyelenggara string `validate:"min=1,max=100" json:"penyelenggara"`
	Judul string `validate:"min=1,max=100" json:"judul"`
	Kota string `validate:"min=1,max=100" json:"kota"`
	Tanggal string `validate:"min=1,max=100" json:"tanggal"`
	Deskripsi string `validate:"min=1,max=100" json:"deskripsi"`
}
