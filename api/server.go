package api

import (
	"context"
	"document"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Server is ...
type Server struct {
	collection *mongo.Collection
	router     *gin.Engine
}

//NewServer is function for initialize server
func NewServer(database string, collection string) Server {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongoDB")

	c := client.Database(database).Collection(collection)
	r := gin.Default()
	s := Server{collection: c, router: r}
	s.router.POST("/register", s.Register)
	s.router.GET("/register/:id", s.GetOneCustomer)
	s.router.GET("/register", s.ListCustomer)
	s.router.POST("/test", s.RegisterApplicant)

	return s
}

//Run is function for run rounter.
func (server Server) Run() {
	server.router.Run()
}

//GetOneCustomer is function that will return JSON
func (server Server) GetOneCustomer(ctx *gin.Context) {
	var req customerRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	filter := bson.D{{Key: "customer_id", Value: req.CustomerID}}
	var result customer
	err = server.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}

//ListCustomer is function
func (server Server) ListCustomer(ctx *gin.Context) {
	var req listCustomerRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var customers []customer
	customers = searchCustomer(server.collection, req.PageSize)
	ctx.JSON(http.StatusOK, customers)

}

//Register is ...
func (server Server) Register(ctx *gin.Context) {
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

func (server Server) RegisterApplicant(ctx *gin.Context) {
	req := document.CreateNewApplicant()
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
