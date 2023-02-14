package servers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Http server wrapper. Incapsulates common logic
type HttpServer struct {
	options HttpOptions
	server  http.Server
	stop    context.CancelFunc
}

// Options
type HttpOptions struct {
	Port     string
	Register func(*mux.Router)
}

// Instantiate new http server
func NewHttpServer(options HttpOptions) *HttpServer {
	listenAddress := ":" + options.Port

	httpServer := http.Server{
		Addr: listenAddress,
	}

	return &HttpServer{
		options,
		httpServer,
		nil,
	}
}

// Start server
func (h *HttpServer) Start(ctx context.Context) {
	router := mux.NewRouter()

	if h.options.Register == nil {
		log.Fatal("http: wrong registration function")
	}

	h.options.Register(router)

	h.server.Handler = router

	ctx, h.stop = context.WithCancel(ctx)

	go onShutdown(ctx, func() {
		if err := h.server.Shutdown(ctx); err != nil {
			log.Println("http: server shutdown error")
		}
	})

	log.Printf("http: server is listening at %v", h.options.Port)
	err := h.server.ListenAndServe()

	if err != nil {
		log.Println("http: server failed: ", err)
	}
}

// Stop server
func (h *HttpServer) Stop() {
	if h.stop != nil {
		h.stop()
	}
}

// Return http json response
func JsonOK(w http.ResponseWriter, data any) {
	bytes, err := json.Marshal(data)

	if err != nil {
		JsonNotOK(w, err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// Return http json error response
func JsonNotOK(w http.ResponseWriter, err error) {
	type errorMessage struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	bytes, _ := json.Marshal(errorMessage{
		Type:    "error",
		Message: err.Error(),
	})

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bytes)
}
