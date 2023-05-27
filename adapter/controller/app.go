package controller

type AppController struct {
	UploadController interface{ UploadController }
}
