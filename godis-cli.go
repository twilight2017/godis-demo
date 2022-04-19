package demo

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	IPPort := "127.0.0.1:9736"

	reader := bufio.NewReader(os.Stdin)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", IPPort)

	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close() //defer关键字用于资源释放，有多个defer表达式时，调用顺序同栈

	for {
		fmt.Print(IPPort + "> ")
		text, _ := reader.ReadString('\n')
		//清除掉回车换行符
		text = strings.Replace(text, "\n", "", -1)
		send2Server(text, conn)

		buff := make([]byte, 1024) //内存分配原语，创建长度为1024的int型slice
		n, err := conn.Read(buff)
		checkError(err)
		if n == 0 {
			fmt.Println(IPPort+"> ", "nil")
		} else {
			fmt.Println(IPPort+"> ", string(buff))
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
