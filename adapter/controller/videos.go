package controller

import (
	"context"
	"github/yogabagas/video-stream-be/domain/service"
	"github/yogabagas/video-stream-be/service/videos/usecase"
)

type VideosControllerImpl struct {
	videosSvc usecase.VideosService
}

type VideosController interface {
	Upload(ctx context.Context, req service.UploadVideosReq) error
	Watch(ctx context.Context, req service.WatchVideosReq) (resp service.WatchVideosResp, err error)
}

func NewVideosController(videosSvc usecase.VideosService) VideosController {
	return &VideosControllerImpl{
		videosSvc: videosSvc,
	}
}

func (uc *VideosControllerImpl) Upload(ctx context.Context, req service.UploadVideosReq) error {
	return uc.videosSvc.Upload(ctx, req)
}

func (uc *VideosControllerImpl) Watch(ctx context.Context, req service.WatchVideosReq) (resp service.WatchVideosResp, err error) {
	return uc.videosSvc.Watch(ctx, req)
}
