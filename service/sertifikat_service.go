package service

import (
	"github.com/Wintec-Yuda/print-certificate.git/model/web"
	"context"
)

type SertifikatService interface {
	Create(ctx context.Context, request web.SertifikatCreateRequest) web.SertifikatResponse
	Update(ctx context.Context, request web.SertifikatUpdateRequest) web.SertifikatResponse
	Delete(ctx context.Context, sertifikatId int)
	FindById(ctx context.Context, sertifikatId int) web.SertifikatResponse
	FindAll(ctx context.Context) []web.SertifikatResponse
}
