package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
