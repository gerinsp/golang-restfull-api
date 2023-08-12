package main

import (
	"belajar-golang-restfull-api/app"
	"belajar-golang-restfull-api/controller"
	"belajar-golang-restfull-api/helper"
	"belajar-golang-restfull-api/middleware"
	"belajar-golang-restfull-api/repository"
	"belajar-golang-restfull-api/service"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	handler := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	fmt.Println("Website berjalan di http://localhost:3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
