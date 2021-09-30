package web

type SertifikatCreateRequest struct {
	Nama string `validate:"required,min=1,max=100" json:"nama"`
	Penyelenggara string `validate:"min=1,max=100" json:"penyelenggara"`
	Judul string `validate:"min=1,max=100" json:"judul"`
	Kota string `validate:"min=1,max=100" json:"kota"`
	Tanggal string `validate:"min=1,max=100" json:"tanggal"`
	Deskripsi string `validate:"min=1,max=100" json:"deskripsi"`
}
