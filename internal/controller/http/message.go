package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"yukiteru-amano/internal/model"
	"yukiteru-amano/internal/service"
)

type messageController struct {
	mm service.MethodMessage
}

func newMessageController(mm service.MethodMessage) *messageController {
	return &messageController{mm: mm}
}

func NewMessageRoutes(r *mux.Router, mm service.MethodMessage) {
	ctrl := newMessageController(mm)

	r.HandleFunc("/log/{id:[0-9]+}", ctrl.logGet).Methods(http.MethodGet)
	r.HandleFunc("/log", ctrl.log).Methods(http.MethodGet, http.MethodPost)
}

func (mc *messageController) log(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body model.Message
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Println(err)
		}

		mc.mm.Set(&body)

		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodGet {
		msgs := mc.mm.GetAll()

		marsh, err := json.Marshal(msgs)
		if err != nil {
			log.Println(err)
		}

		w.Write(marsh)
	}
}

func (mc *messageController) logGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	atoi, _ := strconv.Atoi(vars["id"])

	msg := mc.mm.GetByID(atoi)

	marsh, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	w.Write(marsh)
}
