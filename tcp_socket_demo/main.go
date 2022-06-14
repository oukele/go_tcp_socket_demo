package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tcp_socket_demo/client"
	"tcp_socket_demo/server"
)

func main() {
	client := client.NewClient("127.0.0.1:8088")
	client.StartClient()
	for {
		fmt.Printf("请输入要发送的内容(回车即可发送,退出请输入exit): ")
		// 标准输入
		reader := bufio.NewReader(os.Stdin)
		// 回车代表输入结束
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("终端输入出现错误,请重新输入, err = %v\n", err)
			continue
		}
		if strings.TrimSpace(readString) == "exit" {
			break
		}
		client.SendMessage(readString)
	}
}

func startServer() {
	// 得到一个服务器端的结构体
	server := server.NewServer("")
	// 调用启动的方法
	server.StartServer()
}
