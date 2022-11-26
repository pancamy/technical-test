package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"yt-users-service/helper"
	"yt-users-service/model/web"
)

type JobListServiceImpl struct {
}

func NewJobListServiceImpl() JobListService {
	return &JobListServiceImpl{}
}

func (service *JobListServiceImpl) FindAll(ctx context.Context) []web.JobListResponse {
	response, err := http.Get("http://dev3.dansmultipro.co.id/api/recruitment/positions.json")
	helper.PanicError(err)

	responseData, err := ioutil.ReadAll(response.Body)
	helper.PanicError(err)
	responses := string(responseData)

	var webResponse []web.JobListResponse
	var jsonWebResponse = []byte(responses)

	err = json.Unmarshal(jsonWebResponse, &webResponse)
	helper.PanicError(err)

	return webResponse
}
