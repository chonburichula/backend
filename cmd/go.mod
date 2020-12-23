module main

require (
	api v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.4 // indirect
	middleware v0.0.0
	mongostruct v0.0.0
	myauthorization v0.0.0
)

replace myauthorization => ../myauthorization

replace middleware => ../middleware

replace api => ../api

replace mongostruct => ../mongostruct

go 1.15
