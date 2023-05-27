package controller

import (
	"context"
	"github/yogabagas/video-stream-be/domain/service"
	"github/yogabagas/video-stream-be/service/upload/usecase"
)

type UploadControllerImpl struct {
	uploadSvc usecase.UploadService
}

type UploadController interface {
	Upload(ctx context.Context, req service.UploadVideoReq) error
}

func NewUploadController(uploadSvc usecase.UploadService) UploadController {
	return &UploadControllerImpl{
		uploadSvc: uploadSvc,
	}
}

func (uc *UploadControllerImpl) Upload(ctx context.Context, req service.UploadVideoReq) error {
	return uc.uploadSvc.Upload(ctx, req)
}
