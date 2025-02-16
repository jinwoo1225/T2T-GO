package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	Title     string             `json:"title" validate:"required"`
	Created   string             `json:"created" validate:"required"`
	Creator   primitive.ObjectID `json:"creator" validate:"required"`
	Questions []Question         `json:"questions" validate:"required"`
}

type Question struct {
	Prompt  string   `json:"prompt" validate:"required"`
	Answers []string `json:"answers" validate:"required"`
	Correct int      `json:"correct" validate:"required"`
}
