module learn.go

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.0-00010101000000-000000000000
	github.com/containerd/containerd v1.6.2 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.14+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/json-iterator/go v1.1.12
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/spf13/cobra v1.3.0
	go.etcd.io/etcd/client/v3 v3.5.2
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
	gotest.tools/v3 v3.1.0 // indirect
)

replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/armstrongli/cobra v0.0.0-20211214182251-178edbb247f3 // 原来是 master (branch name)，tidy之后转为精确定位
)
