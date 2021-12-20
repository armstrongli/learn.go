module learn.go

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.3.0
)

replace (
	github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/armstrongli/cobra v0.0.0-20211214182251-178edbb247f3 // 原来是 master (branch name)，tidy之后转为精确定位
)
