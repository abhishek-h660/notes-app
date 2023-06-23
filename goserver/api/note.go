package api

import (
	"encoding/json"
	"net/http"
	"notes/server/config"
	"notes/server/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteController struct {
	*mongo.Database
}

func NewNoteController(db *mongo.Database) *NoteController {
	c := &NoteController{
		Database: db,
	}
	return c
}

func (c *NoteController) AddNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note model.Note
		json.NewDecoder(r.Body).Decode(&note)
		note.Id = primitive.NewObjectIDFromTimestamp(time.Now())
		res, err := note.Save(*c.Database)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func (c *NoteController) DeleteNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//delete
	}
}
func (c *NoteController) ListNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cur, err := c.Database.Collection("notes").Find(config.DefaultCtx(), bson.M{})
		w.Header().Set("content-type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(err)
			return
		}

		var notes []model.Note
		cur.All(config.DefaultCtx(), &notes)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notes)
	}
}
