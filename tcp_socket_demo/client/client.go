package client

import (
	"fmt"
	"net"
	"time"
)

type client struct {
	// 网络协议类型
	networkType string
	// ip地址 + 端口
	ipAddress string
	// 当前的链接
	conn net.Conn
}

// NewClient 通过此函数可以得到一个 client 结构体
func NewClient(ipAddress string) client {
	return client{
		networkType: "tcp",
		ipAddress:   ipAddress,
	}
}

// 获取当前运行的时间
func nowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// StartClient 客户端启动方法
func (c *client) StartClient() {
	conn, err := net.Dial(c.networkType, c.ipAddress)
	if err != nil {
		fmt.Printf("链接服务器端: %v ,time = %v ,出现异常 err = %v\n", c.ipAddress, nowTime(), err)
		return
	}
	fmt.Printf("[%v]链接服务器端:[%v]成功\n", nowTime(), conn.RemoteAddr().String())
	c.conn = conn
}

// SendMessage 发送消息
func (c *client) SendMessage(msg string) {
	writeLen, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("[%v]客户端发送消息至服务器失败, err = %v\n", nowTime(), err)
		return
	}
	fmt.Printf("[%v]客户发送长度为[%d]-[%v]消息成功\n", nowTime(), writeLen, msg)
}
