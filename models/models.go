package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
	Stocks      int64              `json:"stocks"`
}

type Transaction struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Date      int64              `json:"date"`
	Total     int64              `json:"total"`
	Quantity  int64              `json:"quantity"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Status    string             `json:"status" bson:"status"` // pending, canceled, accepted
}
