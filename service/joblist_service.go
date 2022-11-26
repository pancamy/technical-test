package service

import (
	"context"
	"yt-users-service/model/web"
)

type JobListService interface {
	FindAll(ctx context.Context) []web.JobListResponse
}
