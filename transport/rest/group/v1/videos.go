package v1

import (
	"github/yogabagas/video-stream-be/transport/rest/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewVideosV1(v1 *mux.Router, h handler.HandlerImpl) {

	prefix := v1.PathPrefix("/videos").Subrouter()

	prefix.HandleFunc("/upload", h.Upload).Methods(http.MethodPost)
	prefix.HandleFunc("/watch", h.Watch).Methods(http.MethodGet)
}
