package models

import (
	"go.mongo.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID          primitive.ObjectID `bson:"_id"`
	Task        *string            `json:"task"`
	Description *string            `json:"desc"`
	Deadline    *string            `json:"deadline"`
	AssignedTo  *string            `json:"assignedTo"`
}
