package api

import (
	"fmt"
	"net/http"
	"notes/server/config"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	*mux.Router
}

func InitDatabase(db *mongo.Database) {
}

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Dylit Server - 05/04/2023")
	}
}

func NewServer() *Server {
	db := config.MongoClient()

	s := &Server{
		Router: mux.NewRouter(),
	}

	InitDatabase(db)
	notes := NewNoteController(db)
	s.HandleFunc("/api/v1/add_notes", notes.AddNote()).Methods("POST")
	s.HandleFunc("/api/v1/list_notes", notes.ListNotes()).Methods("GET")

	return s

}
