package middleware

import (
	"belajar-golang-restfull-api/helper"
	"belajar-golang-restfull-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writter, request)
	} else {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writter, webResponse)
	}
}
