module github.com/Solar-2020/Account-Backend

go 1.14

require (
	github.com/Solar-2020/Authorization-Backend v1.0.0
	github.com/Solar-2020/GoUtils v1.0.3
	github.com/buaazp/fasthttprouter v0.1.1
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.20.0
	github.com/valyala/fasthttp v1.16.0
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

// replace github.com/Solar-2020/GoUtils => ../GoUtils

// replace github.com/Solar-2020/Authorization-Backend => ../Authorization-Backend
