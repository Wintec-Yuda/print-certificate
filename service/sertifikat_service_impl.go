package service

import (
	"github.com/Wintec-Yuda/print-certificate.git/exception"
	"github.com/Wintec-Yuda/print-certificate.git/helper"
	"github.com/Wintec-Yuda/print-certificate.git/model/domain"
	"github.com/Wintec-Yuda/print-certificate.git/model/web"
	"github.com/Wintec-Yuda/print-certificate.git/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type SertifikatServiceImpl struct {
	SertifikatRepository repository.SertifikatRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewSertifikatService(SertifikatRepository repository.SertifikatRepository, DB *sql.DB, validate *validator.Validate) SertifikatService {
	return &SertifikatServiceImpl{
		SertifikatRepository: SertifikatRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *SertifikatServiceImpl) Create(ctx context.Context, request web.SertifikatCreateRequest) web.SertifikatResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sertifikat := domain.Sertifikat{
		Nama: request.Nama,
		Penyelenggara: request.Penyelenggara,
		Judul: request.Judul,
		Kota: request.Kota,
		Tanggal: request.Tanggal,
		Deskripsi: request.Deskripsi,
	}

	sertifikat = service.SertifikatRepository.Save(ctx, tx, sertifikat)

	return helper.ToSertifikatResponse(sertifikat)
}

func (service *SertifikatServiceImpl) Update(ctx context.Context, request web.SertifikatUpdateRequest) web.SertifikatResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sertifikat, err := service.SertifikatRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	sertifikat.Nama = request.Nama
	sertifikat.Penyelenggara = request.Nama
	sertifikat.Judul = request.Judul
	sertifikat.Kota = request.Kota
	sertifikat.Tanggal = request.Tanggal
	sertifikat.Deskripsi = request.Deskripsi

	sertifikat = service.SertifikatRepository.Update(ctx, tx, sertifikat)

	return helper.ToSertifikatResponse(sertifikat)
}

func (service *SertifikatServiceImpl) Delete(ctx context.Context, sertifikatId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sertifikat, err := service.SertifikatRepository.FindById(ctx, tx, sertifikatId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SertifikatRepository.Delete(ctx, tx, sertifikat)
}

func (service *SertifikatServiceImpl) FindById(ctx context.Context, sertifikatId int) web.SertifikatResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sertifikat, err := service.SertifikatRepository.FindById(ctx, tx, sertifikatId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSertifikatResponse(sertifikat)
}

func (service *SertifikatServiceImpl) FindAll(ctx context.Context) []web.SertifikatResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	sertifikats := service.SertifikatRepository.FindAll(ctx, tx)

	return helper.ToSertifikatResponses(sertifikats)
}
