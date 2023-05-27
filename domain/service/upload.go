package service

import (
	"mime/multipart"
	"time"
)

type UploadVideoReq struct {
	File      multipart.File
	Name      string
	Buffer    []byte
	Size      byte
	CreatedAt time.Time
}
