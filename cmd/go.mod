module main

require (
	api v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.4.4 // indirect
)

replace api => ../api

go 1.15