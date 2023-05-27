package handler

import (
	"encoding/json"
	"github/yogabagas/video-stream-be/domain/service"
	"log"
	"net/http"
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

	req := service.UploadVideoReq{
		File:      file,
		Name:      handler.Filename,
		Size:      byte(handler.Size),
		CreatedAt: time.Now(),
	}

	err = h.Controller.UploadController.Upload(r.Context(), req)
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
