// 模拟服务器端：接收客户端连接，并将客户端发送的内容原样传回客户端

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	// 绑定监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind: %s, satrt listening...", address))

	for {
		//Accecpt 会一直阻塞到有新的连接建立或者Listen中断才会返回
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("Accept err: %v", err))
		}
		go Handle(conn)
		//开启新的go routine处理该连接
	}
}

func Handle(conn net.Conn) {
	// 使用bufio标准库提供的缓冲区功能
	reader := bufio.NewReader(conn)
	for {
		//ReadString会一直阻塞到直到遇到分隔符'\n'
		//遇到分割符后会返回上次遇到的分隔符或者建立连接后收到的所有数据，包括分隔符本身
		//若在遇到分隔符之前遇到异常，ReadString会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			//通常遇到的错误是连接遇到中断或被关闭，用io.EOF表示
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println("err")
			}
			return
		}
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8000")
}
