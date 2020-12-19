module api

go 1.15

require (
	database v0.0.0
	github.com/gin-gonic/gin v1.6.3
	go.mongodb.org/mongo-driver v1.4.4
)

replace database => ../database