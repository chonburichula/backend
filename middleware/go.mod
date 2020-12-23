module middleware

require (
	github.com/gin-gonic/gin v1.6.3
	myauthorization v0.0.0
)

replace myauthorization => ../myauthorization

go 1.15
