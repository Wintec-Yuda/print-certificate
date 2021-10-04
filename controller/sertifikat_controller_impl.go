package controller

import (
	"github.com/Wintec-Yuda/print-certificate.git/helper"
	"github.com/Wintec-Yuda/print-certificate.git/model/web"
	"github.com/Wintec-Yuda/print-certificate.git/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
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
	//sertifikatResponses := controller.SertifikatService.FindAll(request.Context())
	//webResponse := web.WebResponse{
	//	Code:   200,
	//	Status: "OK",
	//	Data:   sertifikatResponses,
	//}
	//
	//helper.WriteToResponseBody(writer, webResponse)

	//data := controller.SertifikatService.FindAll(request.Context())
	//var names = data[0]

	var filepath = path.Join("template", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	//var filepath = path.Join("template", "index.html")
	//var tmpl, err = template.ParseFiles(filepath)
	//if err != nil {
	//	http.Error(writer, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//err = tmpl.Execute(writer, data)
	//if err != nil {
	//	http.Error(writer, err.Error(), http.StatusInternalServerError)
	//}

}

