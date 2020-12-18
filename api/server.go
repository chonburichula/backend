package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	collection *mongo.Collection
	router     *gin.Engine
}

func newServer(collection string) server {

}

func (server server) getOneCustomer(ctx *gin.Context) {
	var req requestCustomer
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	filter := bson.D{{"customer_id", req.CustomerID}}
	var result customer
	err = server.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (server server) listCustomer(ctx *gin.Context) {
	var req listCustomer
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var customers []customer
	customers = searchCustomer(server.collection, req.PageSize)
	ctx.JSON(http.StatusOK, customers)

}

func (server server) Register(ctx *gin.Context) {
	var req customer
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	insertResult, err := server.collection.InsertOne(context.TODO(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, insertResult)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
