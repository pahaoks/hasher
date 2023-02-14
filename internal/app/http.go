package app

import (
	"net/http"

	"example.com/hasher/gen/oapi"
	"example.com/hasher/pkg/servers"
)

// Http handler
type HttpHandler struct {
	hasherService HasherService
}

// New http handler
func NewHttpHandler(
	hasherService HasherService,
) oapi.ServerInterface {
	return &HttpHandler{hasherService}
}

// Get hash
func (h *HttpHandler) GetHash(w http.ResponseWriter, r *http.Request) {
	hash := h.hasherService.GetHash()

	result := oapi.HashResponse{
		Hash:      hash.Value,
		UpdatedAt: hash.CreatedAt,
	}

	servers.JsonOK(w, result)
}
