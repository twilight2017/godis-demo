package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	IPPort := "127.0.0.1:9736"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi Godis")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", IPPort)
	checkError(err)

	//建立连接，如果第二个参数（本地地址）为nil，会自动生成一个本地地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()
	log.Println(tcpAddr, conn.LocalAddr(), conn.RemoteAddr())

	for {
		fmt.Print(IPPort + "> ")
		text, _ := reader.ReadString('\n')
		send2Server(text, conn)

		buf := make([]byte, 20) //make函数生成切片
		n, err := conn.Read(buf)
		checkError(err)
		if n == 0 {
			fmt.Println(IPPort+"> ", "nil")
		} else {
			fmt.Println(IPPort+"> ", string(buf))
		}
	}
}

func send2Server(msg string, conn net.Conn) (n int, err error) {
	data := []byte(msg)
	n, err = conn.Write(data)
	return n, err
}

func checkError(err error) {
	if err != nil {
		log.Println("err ", err.Error())
		os.Exit(1)
	}
}
