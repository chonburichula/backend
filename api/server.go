package api

import (
	"middleware"
	"mongostruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

type Server struct {
	rounter  *gin.Engine
	database string
	redis    *redis.Client
}

//NewServer is function for initialize server
func NewServer(database string, redis *redis.Client) Server {
	r := gin.Default()

	r.POST("/register", register)
	staff := r.Group("/staff")
	staff.Use(middleware.TokenAuthMiddleware())
	staff.GET("/ungraded", getUnGraded)
	staff.GET("/graded", getGraded)
	staff.POST("/update", update)
	server := Server{r, database, redis}
	return server
}

func (server Server) Run() {
	server.rounter.Run()
}

func register(ctx *gin.Context) {
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

func getUnGraded(ctx *gin.Context) {
	ungraded, err := mongostruct.GetUnGradedApplicant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ungraded)
}

func getGraded(ctx *gin.Context) {
	graded, err := mongostruct.GetGradedApplicant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, graded)
}

func update(ctx *gin.Context) {
	var req mongostruct.ApplicantOnlyScore
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	updateResult, err := mongostruct.UpdateStatusAndScore(req.ID, int32(req.Score))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateResult)
}
