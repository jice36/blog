package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jice36/blog/internal/service"
	"github.com/jice36/blog/models"

	"github.com/gorilla/mux"
)

type Server struct {
	config  *ConfigBlog
	Log     *log.Logger
	router  *mux.Router
	service *service.ServiceDB
}

func NerServer(conf *ConfigBlog, s *service.ServiceDB) *Server {
	return &Server{config: conf,
		Log:     log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		router:  mux.NewRouter(),
		service: s,
	}
}

func (s *Server) StartServer() error {
	s.configRouter()

	return http.ListenAndServe(s.config.Server.Host+":"+s.config.Server.Port, s.router)
}

func (s *Server) configRouter() {
	s.router.HandleFunc("/blog/{id}", s.getNote).Methods("GET")
	s.router.HandleFunc("/blog", s.sendNote).Methods("POST")
}

func (s *Server) getNote(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	id, ok := param["id"]
	fmt.Println(id)
	if !ok {
		http.Error(w, "id пользователя не передан", http.StatusNotFound) //todo добавить ошибку и проверить длину id
	}

	resp, err := s.service.GetNotes(id)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) sendNote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}

	data := &models.SendNoteReq{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}

	_, err = s.service.SendNote(data)
	if err != nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
