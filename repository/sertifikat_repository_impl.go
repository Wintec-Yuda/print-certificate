package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Wintec-Yuda/print-certificate.git/helper"
	"github.com/Wintec-Yuda/print-certificate.git/model/domain"
)

type SertifikatRepositoryImpl struct {
}

func NewSertifikatRepository() SertifikatRepository {
	return &SertifikatRepositoryImpl{}
}

func (repository *SertifikatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat) domain.Sertifikat {
	SQL := "insert into sertifikat(nama, penyelenggara, judul, kota, tanggal, deskripsi) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, sertifikat.Nama, sertifikat.Penyelenggara, sertifikat.Judul, sertifikat.Kota, sertifikat.Tanggal, sertifikat.Deskripsi)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	sertifikat.Id = int(id)
	return sertifikat
}

func (repository *SertifikatRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat) domain.Sertifikat {
	SQL := "update Sertifikat set nama = ?, penyelenggara = ?, judul = ?, kota = ?, tanggal = ?, deskripsi = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, sertifikat.Nama, sertifikat.Penyelenggara, sertifikat.Judul, sertifikat.Kota, sertifikat.Tanggal, sertifikat.Deskripsi, sertifikat.Id)
	helper.PanicIfError(err)

	return sertifikat
}

func (repository *SertifikatRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat) {
	SQL := "delete from Sertifikat where id = ?"
	_, err := tx.ExecContext(ctx, SQL, sertifikat.Id)
	helper.PanicIfError(err)
}

func (repository *SertifikatRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, sertifikatId int) (domain.Sertifikat, error) {
	SQL := "select * from Sertifikat where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, sertifikatId)
	helper.PanicIfError(err)
	defer rows.Close()

	sertifikat := domain.Sertifikat{}
	if rows.Next() {
		err := rows.Scan(&sertifikat.Id, &sertifikat.Nama, &sertifikat.Penyelenggara, &sertifikat.Judul, &sertifikat.Kota, &sertifikat.Tanggal, &sertifikat.Deskripsi)
		helper.PanicIfError(err)
		return sertifikat, nil
	} else {
		return sertifikat, errors.New("Sertifikat is not found")
	}
}

func (repository *SertifikatRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Sertifikat {
	SQL := "select * from Sertifikat"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var sertifikats []domain.Sertifikat
	for rows.Next() {
		Sertifikat := domain.Sertifikat{}
		err := rows.Scan(&Sertifikat.Id, &Sertifikat.Nama, &Sertifikat.Penyelenggara, &Sertifikat.Judul, &Sertifikat.Kota, &Sertifikat.Tanggal, &Sertifikat.Deskripsi)
		helper.PanicIfError(err)
		sertifikats = append(sertifikats, Sertifikat)
	}
	return sertifikats
}
