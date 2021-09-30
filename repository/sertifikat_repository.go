package repository

import (
	"context"
	"database/sql"
	"github.com/Wintec-Yuda/print-certificate.git/model/domain"
)

type SertifikatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat) domain.Sertifikat
	Update(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat) domain.Sertifikat
	Delete(ctx context.Context, tx *sql.Tx, sertifikat domain.Sertifikat)
	FindById(ctx context.Context, tx *sql.Tx, sertifikatId int) (domain.Sertifikat, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Sertifikat
}
