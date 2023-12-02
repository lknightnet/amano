package service

import (
	"time"
	"yukiteru-amano/internal/model"
	"yukiteru-amano/pkg/heapofshit"
)

var globalCounter int

func incrementCounter() {
	globalCounter++
}

type messageService struct {
	h *heapofshit.Heap
}

func (m *messageService) Set(message *model.Message) {
	incrementCounter()
	message.ID = globalCounter
	message.Time = time.Now().Format("02.01.2006 15:04:05")

	if message.Text == "" {
		message.Text = "unknown"
	}
	if message.Level == "" {
		message.Level = "unknown"
	}
	if message.MicroServiceName == "" {
		message.MicroServiceName = "unknown"
	}
	m.h.Add(message)
}

func (m *messageService) GetByID(id int) *model.Message {
	elm := m.h.ShowByID(id)
	return elm
}

func (m *messageService) GetAll() []*model.Message {
	elm := m.h.ShowAll()
	return elm
}

func newMessageService(h *heapofshit.Heap) *messageService {
	return &messageService{h: h}
}

var _ MethodMessage = (*messageService)(nil)
