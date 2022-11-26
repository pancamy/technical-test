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

func (controller *JobListControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jobListId := params.ByName("jobListId")

	response := controller.JobList.FindById(request.Context(), jobListId)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *JobListControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jobListCreateRequest := web.JobListCreateRequest{}

	jobListCreateRequest.Description = request.FormValue("description")
	jobListCreateRequest.Location = request.FormValue("location")
	jobListCreateRequest.FullTime = request.FormValue("full_time")
	jobListCreateRequest.Page = request.FormValue("page")

	response := controller.JobList.FindAll(request.Context(), jobListCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}
