package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Modality    *string            `json:"modality"`
}
