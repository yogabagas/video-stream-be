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

type UploadServiceImpl struct{}

type UploadService interface {
	Upload(ctx context.Context, req service.UploadVideoReq) error
}

func NewUploadService() UploadService {
	return &UploadServiceImpl{}
}

func (us *UploadServiceImpl) Upload(ctx context.Context, req service.UploadVideoReq) error {

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

	chunkVideoFile(fileLocation, chunkSize)

	// buffer := make([]byte, chunkSize)

	// for i := 0; ; i++ {

	// 	n, err := req.File.Read(buffer)
	// 	if err != nil && err == io.EOF {
	// 		return err
	// 	}

	// 	if n == 0 {
	// 		break
	// 	}

	// 	chunkFilename := fmt.Sprintf("%s_chunk%d", req.Name, i)
	// 	chunkPath := filepath.Join(dir, "files", chunkFilename)

	// 	chunkFile, err := os.OpenFile(chunkPath, os.O_WRONLY|os.O_CREATE, 0666)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer chunkFile.Close()

	// 	_, err = chunkFile.Write(req.Buffer[:n])
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	// fileLocation := filepath.Join(dir, "files", req.Name)

	// targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer targetFile.Close()

	// if _, err = io.Copy(targetFile, req.File); err != nil {
	// 	return err
	// }

	return nil

}

func chunkVideoFile(inputFile string, chunkSize int) ([]string, error) {
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
