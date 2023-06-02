package service

import (
	"mime/multipart"
	"time"
)

type UploadVideosReq struct {
	File      multipart.File
	Name      string
	Buffer    []byte
	Size      byte
	CreatedAt time.Time
}

type WatchVideosReq struct {
	Range int64
	Name  string
}

type WatchVideosResp struct {
	Size  int64
	Video []byte
}
