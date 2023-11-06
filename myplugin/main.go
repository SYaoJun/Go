// main.go

package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println("无法加载插件:", err)
		return
	}

	sym, err := p.Lookup("MyPluginFunc")
	if err != nil {
		fmt.Println("无法查找插件符号:", err)
		return
	}

	// 调用插件函数
	pluginFunc, ok := sym.(func())
	if !ok {
		fmt.Println("插件符号不是一个函数")
		return
	}

	pluginFunc()
}
