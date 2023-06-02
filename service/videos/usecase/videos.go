package usecase

import (
	"context"
	"fmt"
	"github/yogabagas/video-stream-be/domain/service"
	"io"
	"os"
	"path/filepath"
)

const (
	chunkSize = 1 * 1024 * 1024
)

var (
	chunkCount int
)

type VideosServiceImpl struct{}

type VideosService interface {
	Upload(ctx context.Context, req service.UploadVideosReq) error
	Watch(ctx context.Context, req service.WatchVideosReq) (resp service.WatchVideosResp, err error)
}

func NewVideosService() VideosService {
	return &VideosServiceImpl{}
}

func (vs *VideosServiceImpl) Upload(ctx context.Context, req service.UploadVideosReq) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileLocation := filepath.Join(dir, "files", req.Name)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, req.File); err != nil {
		return err
	}

	go vs.chunkVideoFile(fileLocation, chunkSize)

	return nil

}

func (vs *VideosServiceImpl) Watch(ctx context.Context, req service.WatchVideosReq) (resp service.WatchVideosResp, err error) {

	dir, err := os.Getwd()
	if err != nil {
		return resp, err
	}
	filePath := filepath.Join(dir, "files", req.Name)

	file, err := os.Open(filePath)
	if err != nil {
		return resp, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return resp, err
	}

	resp.Size = fileInfo.Size()

	_, err = file.Seek(25, 0)
	if err != nil {
		return resp, err
	}

	buffer := make([]byte, chunkSize)

	for i := 0; i < chunkCount; i++ {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		resp.Video = buffer[n:]
	}

	return resp, nil
}

func (vs *VideosServiceImpl) chunkVideoFile(inputFile string, chunkSize int) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	chunkCount := int(fileSize) / chunkSize
	if int(fileSize)%chunkSize != 0 {
		chunkCount++
	}

	chunkPaths := make([]string, chunkCount)

	for i := 0; i < chunkCount; i++ {
		chunkPath := fmt.Sprintf("%s.%d", inputFile, i)
		chunkPaths[i] = chunkPath

		chunkFile, err := os.Create(chunkPath)
		if err != nil {
			return nil, err
		}

		// Read and write the chunk of data
		buffer := make([]byte, chunkSize)
		_, err = io.CopyBuffer(chunkFile, file, buffer)
		if err != nil {
			chunkFile.Close()
			return nil, err
		}

		chunkFile.Close()
	}

	return chunkPaths, nil
}
