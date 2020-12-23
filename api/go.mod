module api

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v7 v7.4.0
	go.mongodb.org/mongo-driver v1.4.4
	middleware v0.0.0
	mongostruct v0.0.0
	myauthorization v0.0.0
)

replace myauthorization => ../myauthorization
replace middleware => ../middleware
replace mongostruct => ../mongostruct
