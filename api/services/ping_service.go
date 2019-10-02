package services

import (
	"dogego/api/serializer"
)

type PingService struct {
}

func (service *PingService) Ping() *serializer.Response {
	return &serializer.Response{
		Code:    serializer.CodeOK,
		Message: "Pong",
	}
}
