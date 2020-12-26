module backend/middleware

require (
	backend/myauthorization v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible // indirect
)

replace backend/myauthorization => ../myauthorization

go 1.15
