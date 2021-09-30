package helper

import (
	"github.com/Wintec-Yuda/print-certificate.git/model/domain"
	"github.com/Wintec-Yuda/print-certificate.git/model/web"
)

func ToSertifikatResponse(Sertifikat domain.Sertifikat) web.SertifikatResponse {
	return web.SertifikatResponse{
		Id:   Sertifikat.Id,
		Nama: Sertifikat.Nama,
		Penyelenggara: Sertifikat.Penyelenggara,
		Judul: Sertifikat.Judul,
		Kota: Sertifikat.Kota,
		Tanggal: Sertifikat.Tanggal,
		Deskripsi: Sertifikat.Deskripsi,
	}
}

func ToSertifikatResponses(sertifikats []domain.Sertifikat) []web.SertifikatResponse {
	var sertifikatResponses []web.SertifikatResponse
	for _, sertifikat := range sertifikats {
		sertifikatResponses = append(sertifikatResponses, ToSertifikatResponse(sertifikat))
	}
	return sertifikatResponses
}
