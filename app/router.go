package app

import (
	"github.com/Wintec-Yuda/print-certificate.git/controller"
	"github.com/Wintec-Yuda/print-certificate.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(sertifikatController controller.SertifikatController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/sertifikat", sertifikatController.FindAll)
	router.GET("/api/sertifikat/:sertifikatId", sertifikatController.FindById)
	router.POST("/api/sertifikat", sertifikatController.Create)
	router.PUT("/api/sertifikat/:sertifikatId", sertifikatController.Update)
	router.DELETE("/api/sertifikat/:sertifikatId", sertifikatController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
