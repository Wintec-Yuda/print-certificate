package controller

import (
	"github.com/Wintec-Yuda/print-certificate.git/helper"
	"github.com/Wintec-Yuda/print-certificate.git/model/web"
	"github.com/Wintec-Yuda/print-certificate.git/service"
	"github.com/julienschmidt/httprouter"
	"github.com/thedevsaddam/renderer"
	"net/http"
	"strconv"
)

type SertifikatControllerImpl struct {
	SertifikatService service.SertifikatService
}

func NewSertifikatController(SertifikatService service.SertifikatService) SertifikatController {
	return &SertifikatControllerImpl{
		SertifikatService: SertifikatService,
	}
}

func (controller *SertifikatControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sertifikatCreateRequest := web.SertifikatCreateRequest{}
	helper.ReadFromRequestBody(request, &sertifikatCreateRequest)

	sertifikatResponse := controller.SertifikatService.Create(request.Context(), sertifikatCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   sertifikatResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SertifikatControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sertifikatUpdateRequest := web.SertifikatUpdateRequest{}
	helper.ReadFromRequestBody(request, &sertifikatUpdateRequest)

	sertifikatId := params.ByName("sertifikatId")
	id, err := strconv.Atoi(sertifikatId)
	helper.PanicIfError(err)

	sertifikatUpdateRequest.Id = id

	sertifikatResponse := controller.SertifikatService.Update(request.Context(), sertifikatUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   sertifikatResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SertifikatControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sertifikatId := params.ByName("sertifikatId")
	id, err := strconv.Atoi(sertifikatId)
	helper.PanicIfError(err)

	controller.SertifikatService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SertifikatControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sertifikatId := params.ByName("sertifikatId")
	id, err := strconv.Atoi(sertifikatId)
	helper.PanicIfError(err)

	sertifikatResponse := controller.SertifikatService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   sertifikatResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SertifikatControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	sertifikatResponses := controller.SertifikatService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   sertifikatResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)

	rnd := renderer.New()
	rnd.HTML(writer, http.StatusOK, "template/find_all", sertifikatResponses)
}
