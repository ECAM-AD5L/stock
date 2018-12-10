package schema

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Stock struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}
