package service

import (
	"context"
	"yt-users-service/model/web"
)

type JobListService interface {
	FindById(ctx context.Context, jobListId string) web.JobListResponse
	FindAll(ctx context.Context, request web.JobListCreateRequest) []web.JobListResponse
}
