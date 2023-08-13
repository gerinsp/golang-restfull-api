//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-restfull-api/app"
	"belajar-golang-restfull-api/controller"
	"belajar-golang-restfull-api/middleware"
	"belajar-golang-restfull-api/repository"
	"belajar-golang-restfull-api/service"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var router = wire.NewSet(
	app.NewRouter,
	wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		router,
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
