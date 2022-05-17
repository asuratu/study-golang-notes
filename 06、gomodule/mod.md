## 依赖管理 go mod
1. 生成go mod 文件
```
go mod init [mod name]
```
2. 清理go mod文件
```
go mod tidy
```
3. 同步go mod依赖
```
// 编译当前目录及其子目录
go build ./...
go install ./...
go env GOPATH
// 编译的结果在 /Users/asura/go/bin 目录下
```
