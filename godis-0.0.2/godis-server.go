package main

import (
	"core"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//服务端实例
var godis = new(core.Server)

func main() {
	/*命令行参数处理*/
	argv := os.Args
	argc := len(argv)
	if argc >= 2 {
		/*handle special options*/
		if argv[1] == "-v" || argv[1] == "--version" {
			version()
		}
		if argv[1] == "-h" || argv[1] == "--help" {
			usage()
		}
	}

	/*监听信号 平滑退出*/
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go sigHandler(c)

	/*初始化服务端实例*/
	initServer()

	/*网络处理*/
	netListen, err := net.Listen("tcp", "127.0.0.1:9736")
	if err != nil {
		log.Print("listen err ")
	}
	//checkError(err)
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()

		if err != nil {
			continue
		}
		go handle(conn)
	}
}

//处理请求
func handle(conn net.Conn) {
	c := godis.CreateClient(conn)
	for {
		err := c.ReadQueryFromClient(conn)

		if err != nil {
			log.Println("readQueryFromClient err", err)
			return
		}
		c.ProcessInputBuffer()
		godis.ProcessCommand(c)
		responseConn(conn, c)
	}
}

//响应返回给客户端
func responseConn(conn net.Conn, c *core.Client) {
	conn.Write([]byte(c.Buf))
}

//初始化服务端实例
func initServer() {
	godis.Pid = os.Getpid()
	godis.DbNum = 16
	initDb()
	godis.Start = time.Now().UnixNano() / 1000000

	getCommand := &core.GodisCommand{Name: "get", Proc: core.GetCommand}
	setCommand := &core.GodisCommand{Name: "set", Proc: core.SetCommand}

	godis.Commands = map[string]*core.GodisCommand{
		"get": getCommand,
		"set": setCommand,
	}
}

//初始化db
func initDb() {
	godis.Db = make([]*core.GodisDb, godis.DbNum)
	for i := 0; i < godis.DbNum; i++ {
		godis.Db[i] = new(core.GodisDb)
		godis.Db[i].Dict = make(map[string]*core.GodisObject, 100)
	}
	fmt.Println("init db begin -->", godis.Db)
}

func sigHandler(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			exitHandler()
		default:
			fmt.Println("signal", s)
		}
	}
}

func exitHandler() {
	fmt.Println("exiting smoothly...")
	fmt.Println("bye")
	os.Exit(0)
}

func version() {
	fmt.Println("Godis server v=0.0.2 sha=xxxxxxx:001 malloc=libc-go bits=64")
	os.Exit(0)
}

func usage() {
	println("Usage: ./godis-server [/path/to/redis.conf] [options]")
	println("       ./godis-server - (read config from stdin)")
	println("       ./godis-server -v or --version")
	println("       ./godis-server -h or --help")
	println("Examples:")
	println("       ./godis-server (run the server with default conf)")
	println("       ./godis-server /etc/redis/6379.conf")
	println("       ./godis-server --port 7777")
	println("       ./godis-server --port 7777 --slaveof 127.0.0.1 8888")
	println("       ./godis-server /etc/myredis.conf --loglevel verbose")
	println("Sentinel mode:")
	println("       ./godis-server /etc/sentinel.conf --sentinel")
	os.Exit(0)
}
