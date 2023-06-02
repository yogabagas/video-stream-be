package controller

type AppController struct {
	VideosController interface{ VideosController }
}
