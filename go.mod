module github.com/EDDYCJY/go-gin-example

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.61.0 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200908134130-d2e65c121b96 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /home/rock/go-application/go-gin-example/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => /home/rock/go-application/go-gin-example/middleware
	github.com/EDDYCJY/go-gin-example/models => /home/rock/go-application/go-gin-example/models
	github.com/EDDYCJY/go-gin-example/pkg/setting => /home/rock/go-application/go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/routers => /home/rock/go-application/go-gin-example/routers
)
