package Handler

import (
	requester "github.com/linqcod/interfaces-mirea-sem-5/second_practice/internal/requester/service"
)

type Handler struct {
	service *requester.Service
}

func NewHandler(service *requester.Service) *Handler {
	return &Handler{
		service: service,
	}
}
