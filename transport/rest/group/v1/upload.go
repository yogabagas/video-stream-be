package v1

import (
	"github/yogabagas/video-stream-be/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewUploadV1(v1 *mux.Router, h handler.HandlerImpl) {
	v1.HandleFunc("/upload", h.Upload).Methods(http.MethodPost)
}
