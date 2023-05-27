package registry

import (
	"github/yogabagas/video-stream-be/adapter/controller"
	"github/yogabagas/video-stream-be/service/upload/usecase"
)

func (m *module) NewUploadUsecase() usecase.UploadService {
	return usecase.NewUploadService()
}

func (m *module) NewUploadController() controller.UploadController {
	return controller.NewUploadController(m.NewUploadUsecase())
}
