package registry

import (
	"github/yogabagas/video-stream-be/adapter/controller"
	"github/yogabagas/video-stream-be/service/videos/usecase"
)

func (m *module) NewVideosUsecase() usecase.VideosService {
	return usecase.NewVideosService()
}

func (m *module) NewVideosController() controller.VideosController {
	return controller.NewVideosController(m.NewVideosUsecase())
}
