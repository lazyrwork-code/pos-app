package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TransactionItem struct {
	ProductID   bson.ObjectID `json:"product_id" bson:"product_id"`
	ProductName string        `json:"product_name" bson:"product_name"`
	Price       float64       `json:"price" bson:"price"`
	Quantity    int           `json:"quantity" bson:"quantity"`
	Subtotal    float64       `json:"subtotal" bson:"subtotal"`
}

type Transaction struct {
	ID            bson.ObjectID     `json:"id" bson:"_id,omitempty"`
	Items         []TransactionItem `json:"items" bson:"items"`
	Total         float64           `json:"total" bson:"total"`
	PaymentMethod string            `json:"payment_method" bson:"payment_method"` // cash, transfer
	KasirID       bson.ObjectID     `json:"kasir_id" bson:"kasir_id"`
	KasirName     string            `json:"kasir_name" bson:"kasir_name"`
	CreatedAt     time.Time         `json:"created_at" bson:"created_at"`
}

type TransactionInput struct {
	Items         []TransactionItemInput `json:"items" binding:"required"`
	PaymentMethod string                 `json:"payment_method" binding:"required,oneof=cash transfer"`
}

type TransactionItemInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}
