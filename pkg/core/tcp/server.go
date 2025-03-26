package tcp

import (
	"bufio"
	"fmt"
	"net"
)

// TServer 结构体
type TServer struct {
	Address string
}

// NewServer 创建 TCP 服务器
func NewServer(address string) *TServer {
	return &TServer{Address: address}
}

// Start 启动服务器
func (s *TServer) Start(handler func(net.Conn)) error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on", s.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handler(conn) // 处理连接
	}
}

// 默认处理函数
func DefaultHandler(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)
		conn.Write([]byte("Echo: " + text + "\n"))
	}
}
