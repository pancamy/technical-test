package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JobListController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
