package web

type SertifikatResponse struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
	Penyelenggara string `json:"penyelenggara"`
	Judul string `json:"judul"`
	Kota string `json:"kota"`
	Tanggal string `json:"tanggal"`
	Deskripsi string `json:"deskripsi"`
}
