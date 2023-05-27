package registry

import "github/yogabagas/video-stream-be/adapter/controller"

type module struct{}

type Registry interface {
	NewAppController() controller.AppController
}

type Option func(*module)

func NewRegistry(opts ...Option) Registry {
	m := &module{}

	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *module) NewAppController() controller.AppController {
	return controller.AppController{
		UploadController: m.NewUploadController(),
	}
}
