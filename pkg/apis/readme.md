# 如何生成protobuf go 文件以及生成gorm注解

## 生成protobuf go 文件
### 安装protoc和protobuf-go
下载系统对应的可执行文件并放入 $GOPATH/bin/ 下。
* https://github.com/protocolbuffers/protobuf/releases
* https://github.com/protocolbuffers/protobuf-go/releases

### 在apis文件夹下执行命令
```bash
protoc --go_out=. --plugin= types.proto
```

## 生成gorm注解

### 安装protoc-go-inject-tag

```bash
go install github.com/favadi/protoc-go-inject-tag@latest
```

### 在 apis 文件夹下执行命令

```bash
protoc-go-inject-tag -input="*.pb.go"
```