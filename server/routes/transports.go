package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var transportCollection *mongo.Collection = OpenCollection(Client, "transports")

//add a transport
func AddTransport(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var transport models.Transport

	if err := c.BindJSON(&transport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(transport)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	transport.ID = primitive.NewObjectID()

	result, insertErr := transportCollection.InsertOne(ctx, transport)
	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

//get all transports
func GetTransports(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var transports []bson.M

	cursor, err := transportCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &transports); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	fmt.Println(transports)

	c.JSON(http.StatusOK, transports)
}

//get a transport by its id
func GetTransportById(c *gin.Context) {
	transportID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(transportID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var transport bson.M

	if err := transportCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&transport); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	fmt.Println(transport)

	c.JSON(http.StatusOK, transport)
}

//update transport
func UpdateTransport(c *gin.Context) {
	transID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(transID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var transport models.Transport

	if err := c.BindJSON(&transport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(transport)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	result, err := transportCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"name":        transport.Name,
			"description": transport.Description,
			"modality":    transport.Modality,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

//delete a transport given the id
func DeleteTransport(c *gin.Context) {
	transID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(transID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := transportCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)
}
