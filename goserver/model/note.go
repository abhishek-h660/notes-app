package model

import (
	"notes/server/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
	Text  string             `json:"text" bson:"text"`
}

func (c *Note) Save(db mongo.Database) (*mongo.InsertOneResult, error) {
	return db.Collection("notes").InsertOne(config.DefaultCtx(), c)
}

func (c *Note) Delete(db mongo.Database) (*mongo.DeleteResult, error) {
	return db.Collection("notes").DeleteOne(config.DefaultCtx(), bson.M{"_id": c.Id})
}
