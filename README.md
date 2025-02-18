# Go语言工程实践
## 基本语法
- 切片
- 哈希表
- 字符串
- 函数
- 方法
- 命令行参数
## 并发编程
- 并发
- 通道
- 互斥锁
## 测试
- 插件
- 包管理
- 单元测试
- 覆盖率
- 性能测试
## Web框架
- Gin框架
- RPC
- MySQL

## 注意点
1. 首字母大写的函数才能被其他包引用到。
2. string是不可变的，想要修改string，先转换为[]byte。
3. go的动态库是编译成插件。
4. golang中不存在while循环，使用for循环代替。
5. rune可变，按照unicode输出。byte可变，按照utf-8输出。string不可变，当使用range遍历时，自动转换为[]rune。
6. go语言中的switch-case在进入case语句后直接退出，不需要break，而C/C++语言中会继续向下执行。

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
## 切片
- 可以动态分配大小的连续空间，类似于C++中的vector。
- 切片的传参是传值还是传引用？传引用，修改生效，新增无效。
## golang中的new和make区别
- make是用来初始化map，slice，chan的。
- new是用来创建对象的。
```go
package main

import "fmt"

func main(){
	var a = new(int) // 得到一个int类型的指针
	*a = 10
	fmt.Println(*a)
	var c = new([10]int) // 创建一个数组
	c[0] = 11
	for _, v:= range c{
		fmt.Println(v)
	}
} 

```
## golang的交换
```
s[i], s[j] = s[j], s[i]
```
## 性能测试
```
go test -bench .
```
## 代码覆盖率
```
go test -cover
```