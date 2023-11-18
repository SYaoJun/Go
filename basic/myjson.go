package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 定义一个结构体表示一个简单的用户
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	// 创建一个用户实例
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john.doe@example.com",
	}

	// 创建一个文件，用于写入 JSON 数据
	file, err := os.Create("user.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 创建一个 JSON 编码器，将数据写入文件
	jsonEncoder := json.NewEncoder(file)
	// 其中encode就是写入
	// decode就是取出
	err = jsonEncoder.Encode(user)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("JSON data has been written to user.json")

	var decodedUser User
	file.Seek(0, 0)
	jsonDecoder := json.NewDecoder(file)
	// 使用 Decode 方法将 JSON 数据解码到用户结构体中
	err = jsonDecoder.Decode(&decodedUser)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 打印解码后的用户结构体
	fmt.Println("Decoded User:", decodedUser)
}
