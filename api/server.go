package api

import (
	"mongostruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Server is ...
type Server struct {
	router *gin.Engine
}

//NewServer is function for initialize server
func NewServer() Server {
	r := gin.Default()
	s := Server{router: r}
	s.router.POST("/register", s.register)
	s.router.GET("/listaccount/:status", s.showListApplicant)
	//s.router.GET("/register/:id", s.getOneCustomer)
	//s.router.GET("/register", s.listCustomer)
	return s
}

//Run is ...
func (server Server) Run() {
	server.router.Run()
}

func (server Server) showListApplicant(ctx *gin.Context) {
	req := ctx.Params.ByName("status")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	if req == "ungraded" {
		applicant, err := mongostruct.ShowUnGradedApplicant()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, applicant)
		return
	}
	applicant, err := mongostruct.ShowGradedApplicant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, applicant)
}

// func (server Server) getOneCustomer(ctx *gin.Context) {
// 	var req customerRequest
// 	err := ctx.ShouldBindUri(&req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	filter := bson.D{{Key: "customer_id", Value: req.CustomerID}}
// 	var result customer
// 	err = server.collection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, result)

// }

// func (server Server) listCustomer(ctx *gin.Context) {
// 	var req listCustomerRequest
// 	err := ctx.ShouldBindQuery(&req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	var customers []customer
// 	customers = searchCustomer(server.collection, req.PageSize)
// 	ctx.JSON(http.StatusOK, customers)

// }

func (server Server) register(ctx *gin.Context) {
	var req mongostruct.Applicant
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	insertResult, err := mongostruct.Insert(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, insertResult)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
