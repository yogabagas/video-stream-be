package handler

import (
	"encoding/json"
	"fmt"
	"github/yogabagas/video-stream-be/domain/service"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (h *HandlerImpl) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to parse file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	req := service.UploadVideosReq{
		File:      file,
		Name:      handler.Filename,
		Size:      byte(handler.Size),
		CreatedAt: time.Now(),
	}

	err = h.Controller.VideosController.Upload(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, "error in processing", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode("OK"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HandlerImpl) Watch(w http.ResponseWriter, r *http.Request) {

	// Ensure there is a range given for the video
	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" {
		http.Error(w, "Required request range header", http.StatusBadRequest)
		return
	}

	// Get video stats (about 61MB)
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Failed to get directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(dir, "files", "Dragon Ball Xenoverse 2 - Hero of Justice Pack 2 Launch Trailer _ PS4 Games.mp4")
	videoFile, err := os.Open(filePath)
	if err != nil {
		log.Println("err", err.Error())
		http.Error(w, "Failed to open video file", http.StatusInternalServerError)
		return
	}
	defer videoFile.Close()

	videoFileInfo, err := videoFile.Stat()
	if err != nil {
		http.Error(w, "Failed to get video file info", http.StatusInternalServerError)
		return
	}
	videoSize := videoFileInfo.Size()

	// Parse Range
	// Example: "bytes=32324-"
	chunkSize := 1 * 1024 * 1024 // 10MB
	rangeParts := strings.Split(rangeHeader[:len(rangeHeader)-1], "=")[1]

	rangeStart, _ := strconv.ParseInt(rangeParts, 10, 64)

	rangeEnd := rangeStart + int64(chunkSize) - 1

	if rangeEnd >= videoSize {
		rangeEnd = videoSize - 1
	}

	// Create headers
	contentLength := rangeEnd - rangeStart + 1

	// HTTP Status 206 for Partial Content
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", rangeStart, rangeEnd, videoSize))
	w.Header().Add("Accept-Ranges", "bytes")
	w.Header().Add("Content-Length", strconv.FormatInt(contentLength, 10))

	w.WriteHeader(http.StatusPartialContent)

	// Create video read stream for this particular chunk
	_, err = videoFile.Seek(rangeStart, 0)
	if err != nil {
		http.Error(w, "Failed to seek video file", http.StatusInternalServerError)
		return
	}

	// Stream the video chunk to the client
	buffer := make([]byte, chunkSize)
	for {
		n, err := videoFile.Read(buffer)
		if err != nil {
			break
		}

		if n > 0 {
			_, err = w.Write(buffer[:n])
			if err != nil {
				break
			}

			// Flush the response writer
			w.(http.Flusher).Flush()
		}
	}
}
