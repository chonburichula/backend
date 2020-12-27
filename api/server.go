package api

import (
	"net/http"

	"github.com/chonburichula/backend/middleware"
	"github.com/chonburichula/backend/mongostruct"
	"github.com/chonburichula/backend/myauthorization"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	rounter  *gin.Engine
	database string
}
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

//NewServer is function for initialize server
func NewServer(database string) Server {
	r := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://chulachon.com", "http://www.chulachon.com"}
	// r.Use(cors.New(config))
	r.Use(cors.Default())
	r.POST("/register", register)
	r.POST("/login", login)
	staff := r.Group("/staff")
	staff.Use(middleware.TokenAuthMiddleware())
	staff.POST("/ungraded", getUnGraded)
	staff.POST("/graded", getGraded)
	staff.POST("/update", update)
	server := Server{rounter: r, database: database}
	return server
}

func (server Server) Run(addr string) {
	server.rounter.Run(addr)
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

func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := myauthorization.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := myauthorization.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
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
