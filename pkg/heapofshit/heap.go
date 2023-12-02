package heapofshit

import "yukiteru-amano/internal/model"

type Heap struct {
	shits []*model.Message
}

func NewHeap(capacity int) *Heap {
	return &Heap{shits: make([]*model.Message, 0, capacity)}
}

func (h *Heap) Add(elm *model.Message) {
	h.shits = append(h.shits, elm)
}

func (h *Heap) ShowByID(id int) *model.Message {
	var res *model.Message
	for _, value := range h.shits {
		if value.ID == id {
			res = value
		}
	}
	return res
}

func (h *Heap) ShowAll() []*model.Message {
	return h.shits
}
