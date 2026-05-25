package handlers

import (
	"context"
	"net/http"
	"pos-app/config"
	"pos-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateTransaction(c *gin.Context) {
	var input models.TransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil kasir dari token
	kasirID, _ := bson.ObjectIDFromHex(c.GetString("userID"))
	kasirName, _ := c.Get("userEmail")

	var items []models.TransactionItem
	var total float64

	for _, item := range input.Items {
		productID, err := bson.ObjectIDFromHex(item.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID tidak valid"})
			return
		}

		// Cari produk & cek stok
		var product models.Product
		err = config.DB.Collection("products").FindOne(context.Background(), bson.M{"_id": productID}).Decode(&product)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stok " + product.Name + " tidak cukup"})
			return
		}

		subtotal := product.Price * float64(item.Quantity)
		total += subtotal

		items = append(items, models.TransactionItem{
			ProductID:   productID,
			ProductName: product.Name,
			Price:       product.Price,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

		// Kurangi stok
		_, err = config.DB.Collection("products").UpdateOne(
			context.Background(),
			bson.M{"_id": productID},
			bson.M{"$inc": bson.M{"stock": -item.Quantity}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update stok"})
			return
		}
	}

	transaction := models.Transaction{
		ID:            bson.NewObjectID(),
		Items:         items,
		Total:         total,
		PaymentMethod: input.PaymentMethod,
		KasirID:       kasirID,
		KasirName:     kasirName.(string),
		CreatedAt:     time.Now(),
	}

	_, err := config.DB.Collection("transactions").InsertOne(context.Background(), transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan transaksi"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	cursor, err := config.DB.Collection("transactions").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func GetReport(c *gin.Context) {
	pipeline := bson.A{
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"$dateToString": bson.M{
						"format": "%Y-%m-%d",
						"date":   "$created_at",
					},
				},
				"total_penjualan":  bson.M{"$sum": "$total"},
				"jumlah_transaksi": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{"_id": -1},
		},
	}

	cursor, err := config.DB.Collection("transactions").Aggregate(context.Background(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
