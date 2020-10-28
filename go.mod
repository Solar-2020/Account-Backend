module github.com/Solar-2020/Account-Backend

go 1.14

require (
	github.com/Solar-2020/Authorization-Backend v0.0.0-20201028130607-d15b917ed022 // indirect
	github.com/Solar-2020/GoUtils v0.0.0-20201028130128-34e4f0f5a23d
	github.com/buaazp/fasthttprouter v0.1.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.8.0
	github.com/rs/zerolog v1.20.0
	github.com/valyala/fasthttp v1.16.0
	google.golang.org/grpc v1.27.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/Solar-2020/GoUtils => ../GoUtils

replace github.com/Solar-2020/Authorization-Backend => ../Authorization-Backend
