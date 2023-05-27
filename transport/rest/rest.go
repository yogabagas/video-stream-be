package rest

import (
	"github/yogabagas/video-stream-be/registry"
	groupV1 "github/yogabagas/video-stream-be/transport/rest/group/v1"
	"github/yogabagas/video-stream-be/transport/rest/handler"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Options struct {
	Port         string
	Address      string
	Mux          *mux.Router
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type Handler struct {
	options *Options
}

func NewRest(o *Options) *Handler {
	reg := registry.NewRegistry()
	appController := reg.NewAppController()

	handlerImpl := handler.HandlerImpl{
		Controller: appController,
	}

	r := mux.NewRouter()

	r.Path("/health").HandlerFunc(handlerImpl.HealthCheck)

	v1 := r.PathPrefix("/v1").Subrouter()

	groupV1.NewUploadV1(v1, handlerImpl)

	o.Mux = r

	return &Handler{options: o}
}

func (h *Handler) Serve() {
	log.Printf("server serve at port %s", h.options.Port)
	http.ListenAndServe(h.options.Port, h.options.Mux)
}
