module main

require (
	backend/api v0.0.0
	backend/middleware v0.0.0
	backend/mongostruct v0.0.0
	backend/myauthorization v0.0.0
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.4 // indirect
)

replace backend/myauthorization => ../myauthorization

replace backend/middleware => ../middleware

replace backend/api => ../api

replace backend/mongostruct => ../mongostruct

go 1.15
