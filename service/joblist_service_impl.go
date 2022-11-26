package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yt-users-service/exception"
	"yt-users-service/helper"
	"yt-users-service/model/web"
)

type JobListServiceImpl struct {
}

func NewJobListServiceImpl() JobListService {
	return &JobListServiceImpl{}
}

func (service *JobListServiceImpl) FindById(ctx context.Context, jobListId string) web.JobListResponse {
	response, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions/" + jobListId)
	helper.PanicError(err)

	if response.Body == nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	responseData, err := ioutil.ReadAll(response.Body)
	helper.PanicError(err)
	responses := string(responseData)

	var jobLisResponse web.JobListResponse
	var jsonWebResponse = []byte(responses)

	err = json.Unmarshal(jsonWebResponse, &jobLisResponse)
	helper.PanicError(err)

	return jobLisResponse
}

func (service *JobListServiceImpl) FindAll(ctx context.Context, request web.JobListCreateRequest) []web.JobListResponse {
	response, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json?location=" + request.Location + "&description=" + request.Description + "&page=" + request.Page)
	helper.PanicError(err)

	if response.StatusCode != 200 {
		panic(exception.NewNotFoundError(err.Error()))
	}

	responseData, err := ioutil.ReadAll(response.Body)
	helper.PanicError(err)
	responses := string(responseData)

	var jobLisResponse []web.JobListResponse
	var jsonWebResponse = []byte(responses)

	err = json.Unmarshal(jsonWebResponse, &jobLisResponse)
	helper.PanicError(err)

	return jobLisResponse
}
