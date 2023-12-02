package service

import (
	"yukiteru-amano/internal/model"
	"yukiteru-amano/pkg/heapofshit"
)

type MethodMessage interface {
	Set(message *model.Message)
	GetByID(id int) *model.Message
	GetAll() []*model.Message
}

type Services struct {
	MD MethodMessage
}

type Dependencies struct {
	*heapofshit.Heap
}

func NewServices(deps *Dependencies) *Services {
	return &Services{MD: newMessageService(deps.Heap)}
}
