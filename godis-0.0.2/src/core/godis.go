package core

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

//Client与服务器端建立连接后创建一个Client结构
type Client struct {
	Cmd      *GodisCommand
	Argv     []*GodisObject //Client队列
	Argc     int
	Db       *GodisDb
	QueryBuf string
	Buf      string
}

type GodisCommand struct {
	Name string
	Proc cmdFunc
}

//命令指针函数
type cmdFunc func(c *Client, s *Server)

//Server 服务器端实例结构体
type Server struct {
	Db               []*GodisDb
	DbNum            int
	Start            int64
	Port             int32
	RdbFilename      string
	AofFilename      string
	NextClientID     int32
	SystemMemorySize int32
	Clients          int32
	Pid              int
	Commands         map[string]*GodisCommand
	Dirty            int64
	AofBuf           []string
}

//use map[string]* as type dict
//使用Go原生数据map作为redis中dict结构体，暂时不对dict造轮子
type dict map[string]*GodisObject

//GodisDb db结构体
type GodisDb struct {
	Dict    dict
	Expires dict
	ID      int32
}

//SetCommangd cmd of set
func SetCommand(c *Client, s *Server) {
	objKey := c.Argv[1]
	objValue := c.Argv[2]
	if c.Argc != 3 {
		addReply(c, CreateObject(ObjectTypeString, "(error) ERR wrong number of arguments for 'set' command"))
		return
	}
	if stringKey, ok1 := objKey.Ptr.(string); ok1 {
		if stringValue, ok2 := objValue.Ptr.(string); ok2 {
			c.Db.Dict[stringKey] = CreateObject(ObjectTypeString, stringValue) //每个dict的value字段存储类型和实际value
		}
	}
	addReply(c, CreateObject(ObjectTypeString, "OK"))
}

//GetCommand get命令实现
func GetCommand(c *Client, s *Server) {
	o := lookupKey(c.Db, c.Argv[1])
	if o != nil {
		addReply(c, o)
	} else {
		addReply(c, CreateObject(ObjectTypeString, "nil")) //查不到时创建一个对象,value为空
	}
}

//addReply添加回复
func addReply(c *Client, o *GodisObject) {
	c.Buf = o.Ptr.(string)
}

//ProcessCommand执行命令
func (s *Server) ProcessCommand(c *Client) {
	v := c.Argv[0].Ptr
	name, ok := v.(string)
	if !ok {
		log.Println("error cmd")
		os.Exit(1)
	}
	cmd := lookupCommand(name, s)
	if cmd != nil {
		c.Cmd = cmd
		call(c, s)
	} else {
		addReply(c, CreateObject(ObjectTypeString, fmt.Sprintf("(error) ERR unknown command '%s'", name)))
	}
}

//lookupCommand命令
func lookupCommand(name string, s *Server) *GodisCommand {
	if cmd, ok := s.Commands[name]; ok { //按name从Commands队列查找Command
		return cmd
	}
	return nil
}

//call()真正调用命令
func call(c *Client, s *Server) {
	fmt.Println("Server call", c.Cmd.Name, c.Cmd.Proc, s.Db)
	c.Cmd.Proc(c, s)
}
func lookupKey(db *GodisDb, key *GodisObject) (ret *GodisObject) {
	if o, ok := db.Dict[key.Ptr.(string)]; ok {
		return o
	}
	return nil
}

//CreateClient 连接建立 创建client记录当前连接
func (s *Server) CreateClient(conn net.Conn) (c *Client) {
	c = new(Client)
	c.Db = s.Db[0]
	c.Argv = make([]*GodisObject, 5)
	c.QueryBuf = ""
	return c
}

//ReadQueryFromClient 读取客户端请求信息
func (c *Client) ReadQueryFromClient(conn net.Conn) (err error) {
	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		log.Println("conn.Read err!=nil", err, "---len---", n, conn)
		conn.Close()
		return err
	}
	tmp := string(buff)
	parts := strings.Split(tmp, "\n")
	c.QueryBuf = parts[0]
	return nil
}

//ProcessInputBuffer处理客户端请求信息
func (c *Client) ProcessInputBuffer() {
	r := regexp.MustCompile("[^\\s]+")
	parts := r.FindAllString(strings.Trim(c.QueryBuf, " "), -1)
	argc, argv := len(parts), parts
	c.Argc = argc
	j := 0
	for _, v := range argv {
		c.Argv[j] = CreateObject(ObjectTypeString, v)
		j++
	}
}
