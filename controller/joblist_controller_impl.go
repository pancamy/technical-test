package controller

import (
	"net/http"
	"yt-users-service/helper"
	"yt-users-service/model/web"
	"yt-users-service/service"

	"github.com/julienschmidt/httprouter"
)

type JobListControllerImpl struct {
	JobList service.JobListService
}

func NewJobListControllerImpl(jobList service.JobListService) JobListController {
	return &JobListControllerImpl{
		JobList: jobList,
	}
}

func (controller *JobListControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.JobList.FindAll(request.Context())

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}
