package core

import=(
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

//Client与服务器连接之后即创建一个Client结构
type Client struct{
	Cmd *GodisCommand
	Argv []*GodisObject
	Argc int
	Db *GodisDb
	QueryBuf string
	Buf string
}

//GodisCommand redis命令结构
type GodisCommand struct{
	Name string
	Proc cmdFunc
}

//命令函数指针
type cmdFunc func(c *CLient, s*Server)

//Server 服务器实例结构体
type Server struct{
	Db []*GodisDb
	DbNum int
	Start int64
	Port int32
	RdbFilename string
	AofFilename string
	NextClientID int32
	SystemMemorySize int32
	Clients int32
	Pid int
	Commands map[string]*GodisCommand
	Dirty int64
	AofBuf []string
}

//use map[string]* as type dict
//使用Go原生数据结构map作为redis中的dict结构体 
type dict map[string]*GodisObject

//GodisDb db结构体
type GodisDb struct{
	Dict dict
	Expires dict
	ID int32
}

//SetCommand cmd for set
func SetCommand(c *Client s *Server){
	objKey := c.Argv[1]
	objValue := c.Argv[2]
	if c.Argc != 3{
		addReply(c, CreateObject(ObjectTypeString, "(error) ERR wrong number of arguments for 'set' command"))
		return
	}
	if stringKey, ok1 := objKey.Ptr.(string);ok1{
		if stringValue, ok2 := objValue.Ptr.(string); ok2{
			c.Db.Dict[stringKey] = CreateObject(ObjectTypeString, stringValue)
		}
	}
	addReply(c, CreateObject(ObjectTypeString, "OK"))
}

//GetCommangd get命令实现
func GetCommand(c *Client, s *Server){
	o := looupKey(c.Db, c.Argv[1])
	if o != nil{
		addReply(c, o)
	}else{
		addReply(c, CreateObject(ObjectTypeString, "nil"))
	}
}

//addReply添加回复
func addReply(c *Client, o *GodisObject){
	c.buf = o.Ptr.(string)
}

//ProcessCommand 执行命令
func (s *Server) ProcessCommand(c *Client){
	v := c.Argv[0].Ptr
	name, ok := v.(string)
	if !ok{
		log.Println("err cmd")
		os.Exit(1)_
	}
	cmd := lookupCommand(name, s)
	if cmd != nil{
		c.Cmd = cmd
		call(c, s)
	}else{
		addReply(c, CreateObject(ObjectTypeString, fmt.Sprintf("(error) ERR unkown command '%s'", name)))
	}
}

//lookupCommand查找命令
func lookupCommand(name string, s *Server) *GodisCommand{
	if cmd, ok := s.Commands[name]; ok{
		return cmd
	}
}

//call真正调用命令
func call(c *Client, s *Server){
	//fmt.Println("server call", c.Cmd.Name, c.Cmd.Proc, s.Db)
	c.Cmd.Proc(c, s)
}

//lookupKey查找对应的键
func lookupKey(db *GodisDb, key *GodisObject)(ret *GodisObject){
	if o, ok := db.Dict[key.Ptr.(string)];ok{
		return o
	}
	return nil
}

//CreateClient 建立连接，创建client记录当前连接
func(s *Server) CreateClient(conn net.Conn)(c *Client){
	c = new(Client)
	c.Db = s.Db[0]
	c.Argv = make([]*GodisObject, 5)
	c.QueryBuf =""
	return c
}

//ReadQueryFromClient 读取客户端请求信息
func (c *Client) ReadQueryFromClient(conn net.Conn)(err error){
	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil{
		log.Println("conn.Read err != nil", err, "---len---", n, conn)
		conn.Close()
		return err
	}
	tmp := string(buff)
	parts := string.Split(tmp, "\n")
	c.QueryBuf = parts[0]
	return nil
}

//ProcessInputBuffer 处理客户端请求信息
func (c *Client)ProcessInputBuffer(){
	r := regexp.MustCompile("[^\\s]+")
	parts := r.FindAllString(strings.Trim(c.QueryBuf, " "), -1)
	argc, argv := len(parts), parts
	c.Argc = argc
	j := 0
	for _, v := range argv{
		c.Argv[j] = CreateObject(ObjectTypeString, v)
		j++
	}
}