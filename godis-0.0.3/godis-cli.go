package main

import (
	"bufio"
	"fmt"
	"godis/core/proto"
	"log"
	"net" //Go语言标准库里提供的net包，支持基于IP层、TCP/UDP层及更高层面（如HTTP、FTP、SMTP）的网络操作
	"os"
	"strings"
)

func main(){
	IPPort := "127.0.0.1:9736"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi Godis")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", IPPort)
	checkError(err)

	//建立连接 如果低二个参数（本地地址）为nil，会自动生成一个本地地址
	conn, err = net.DialTCP("tcp", nil, tcpAddr)//Dial函数用于创建网络连接
	checkError(err)
	defer conn.Close()//推迟连接的关闭，先做下面的收尾工作
	
	for{
		fmt.Print(IPPort + "> ")
		text,_ := reader.ReadString('\n')
		//清除掉所有换行回车符
		text = strings.Replace(text, "\n", "", -1)
		send2Server(text, conn)

		buff:=make([]byte, 1024)//申请切片
		n, err:= conn.Read(buff)//成功建立连接后，可以使用conn的Write()和Read()方法读写数据
		resp, er := proto.DecodeFromBytes(buff)
		checkError(err)
        if n==0{
			fmt.Println(IPPort+ "> ", "nil")
		}else if er ==nil{//解码成功
			fmt.Println(IPPort + ">", string(resp.Value))
		}else{
			fmt.Println(IPPort+ "> ", "err server response")
		}
	}
}

func send2Server(msg string, conn net.Conn)(n int, err error){
	p,e := proto.EncodeCmd(msg)
	if e != nil{
		return 0, e
	}
	n, err = conn.Write(p)//将数据写给Server端
	return n, err
}
/*
error类型及其使用
type error interface{
	Error() string
}
其中只声明了一个Error()方法，用于返回字符串类型的错误消息
*/
func checkError(err error){
	if err != nil{
		log.Println("errpr:" err.Error())
		os.Exit(1)//让当前程序以给出的状态码code退出，一般0表示成功，非0表示出错，程序会立刻终止
	}
}