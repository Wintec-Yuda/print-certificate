package main

import (
	"github.com/Wintec-Yuda/print-certificate.git/app"
	"github.com/Wintec-Yuda/print-certificate.git/controller"
	"github.com/Wintec-Yuda/print-certificate.git/helper"
	"github.com/Wintec-Yuda/print-certificate.git/repository"
	"github.com/Wintec-Yuda/print-certificate.git/service"
	"github.com/thedevsaddam/renderer"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

var Rnd *renderer.Render

func init() {
	Rnd = renderer.New(
		renderer.Options{
			ParseGlobPattern: "html/*.html",
		},
	)
}

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewSertifikatRepository()
	categoryService := service.NewSertifikatService(categoryRepository, db, validate)
	categoryController := controller.NewSertifikatController(categoryService)
	router := app.NewRouter(categoryController)

	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: middleware.NewAuthMiddleware(router),
	//}

	err := http.ListenAndServe(":1000", router)
	helper.PanicIfError(err)
}
