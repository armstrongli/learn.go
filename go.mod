module learn.go

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.0-00010101000000-000000000000
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.2
	github.com/json-iterator/go v1.1.12
	github.com/spf13/cobra v1.3.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/armstrongli/cobra v0.0.0-20211214182251-178edbb247f3 // 原来是 master (branch name)，tidy之后转为精确定位
)
