# Go语言工程实践
- 字符串
- 单元测试
- 包管理
- Gin框架
- 并发

## 注意点
1. 首字母大写的函数才能被其他包引用到。
2. string是不可变的，想要修改string，先转换为[]byte。
3. go的动态库是编译成插件。

```sh
go build -buildmode=plugin plugin.go
```
## Go语言常用命令
```sh
# 生成二进制文件
go build main.go 
# 运行测试
go test hello_test.go
# 编译并运行
go run hello.go
```
## Go引用自定义包
- 从module名开始，到包名。
- 实际调用的时候要跟上包名。
```go
import "yaojungo/mypkg"
mypkg.CustomPkgFunc()
```
## Go引用第三方库
```sh
go get -u github.com/gin-gonic/gin
cd mygin
go run main.go
# 服务启动后点击这个链接 http://127.0.0.1:8080/hello
```
