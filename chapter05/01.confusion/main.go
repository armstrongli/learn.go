package main

func main() {
	dd := &downloadFromNetDisk{
		secret:   &mobileTokenDynamic{mobileNumber: "136132456789"},
		filePath: "接口编程.ppt",
	}
	dd.DownloadFile()
}

type DynamicSecret interface {
	GetSecret() string
}

type mobileTokenDynamic struct {
	mobileNumber string
}

func (d *mobileTokenDynamic) GetSecret() string {
	return "something"
}

// 通常开发的时候，第一个版本叫做：happy path
// 剩下的是痛：它无法应对变更。简单的变更会带来更痛苦的维护。
