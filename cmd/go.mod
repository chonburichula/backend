module main

require (
	api v0.0.0
	mongostruct v0.0.0
	middleware v0.0.0
	myauthorization v0.0.0
	github.com/gin-gonic/gin v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.4.4 // indirect
)

replace myauthorization => ../myauthorization
replace middleware => ../middleware
replace api => ../api
replace mongostruct => ../mongostruct
go 1.15
