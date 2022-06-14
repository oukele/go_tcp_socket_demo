package server

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type server struct {
	// 网络协议类型
	networkType string
	// ip地址 + 端口
	ipAddress string
}

// NewServer 通过此函数可以得到一个 server 结构体
func NewServer(ipAddress string) server {
	if ipAddress == "" {
		ipAddress = "127.0.0.1:8088"
	}
	return server{
		// 默认值 -- 方便开发
		networkType: "tcp",
		ipAddress:   ipAddress,
	}
}

// StartServer 启动服务端的方法
func (s server) StartServer() {

	// 1. 表示使用的网络协议：tcp
	// 2. 对本机进行监听 ip为 xxx.xxx.xxx.xxx + 端口为 xxx
	listener, err := net.Listen(s.networkType, s.ipAddress)
	if err != nil {
		panic(fmt.Sprintf("服务器启动失败, err = %v \n", err))
	}

	fmt.Printf("服务器端启动成功, %v , time: %v \n", listener.Addr().String(), nowTime())

	// 服务器不断的接收客户端的信息
	for {
		conn, err := listener.Accept()
		// 客户端的信息
		remoteAddr := conn.RemoteAddr()
		if err != nil {
			fmt.Printf("客户端[%v]: %v 链接服务器端出现异常, err = %v \n", nowTime(), remoteAddr.String(), err)
			// 跳过此客户端链接不在往下执行
			continue
		}
		fmt.Printf("客户端[%v]: %v 链接服务器端成功 \n", nowTime(), remoteAddr.String())
		// 处理客户端的消息..
		// 有多个客户端链接，就有多少个协程
		go process(conn)
	}
}

// 获取当前运行的时间
func nowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 用于处理客户端消息的函数
func process(conn net.Conn) {
	// 客户端断开链接后要及时关闭
	defer conn.Close()
	connRemoteAddr := conn.RemoteAddr().String()
	// 循环读取客户端的消息
	for {
		// 创建一个1024长度的 byte 切片用于存储客户端的消息
		bytes := make([]byte, 1024)
		// 等待客户端通过 conn 发送消息
		// 如果客户端一直没有发送，那么此协程就阻塞在这里
		readLen, err := conn.Read(bytes)
		if err != nil {
			errStr := err.Error()
			contains := strings.Contains(errStr, "An existing connection was forcibly closed by the remote host")
			if contains || errStr == "EOF" {
				fmt.Printf("客户端[%v]: %v 断开链接\n", nowTime(), connRemoteAddr)
				return
			} else {
				fmt.Printf("服务器读取客户端[%v]: %v 消息出现异常, err = %v \n", nowTime(), connRemoteAddr, err)
				return
			}
		}
		// 服务器端显示客户端的消息
		// bytes[:readLen] ==> 只打印有效的数据长度
		fmt.Printf("server[%v]-客户端[%v]说: %v\n", nowTime(), connRemoteAddr, string(bytes[:readLen]))
	}
}
