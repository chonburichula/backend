module backend/api

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v7 v7.4.0
	go.mongodb.org/mongo-driver v1.4.4
	backend/middleware v0.0.0
	backend/mongostruct v0.0.0
	backend/myauthorization v0.0.0
)

replace backend/myauthorization => ../myauthorization
replace backend/middleware => ../middleware
replace backend/mongostruct => ../mongostruct
