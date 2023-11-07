package main

import "fmt"

func main() {
	// 创建切片
	var strList []string
	strList = append(strList, "hello")
	strList = append(strList, "nihao")
	fmt.Println(strList)
	// 创建哈希表与获取哈希表中的值
	mp := make(map[string]int)
	mp["hello"] = 200
	mp["nihao"] = 300
	mp["yj"] = 300
	// 获取值
	x, ok := mp["hello"]
	if ok {
		fmt.Println(x)
	}
	fmt.Println("delete before")
	for k, v := range mp {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
	delete(mp, "hello")
	fmt.Println("delete after")
	for k, v := range mp {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

}
