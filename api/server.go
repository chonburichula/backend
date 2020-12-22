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
	s.router.POST("/register", register)
	s.router.GET("/ungraded", getUnGraded)
	s.router.GET("/graded", getGraded)
	s.router.POST("/update", update)
	return s
}

//Run is ...
func (server Server) Run() {
	server.router.Run()
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
