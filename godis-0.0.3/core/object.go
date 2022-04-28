package core

// GodisObject是对特定类型的数据的包装
type GodisObject struct {
	ObjectType int
	Ptr        interface{}
}

const ObjectTypeString = 1

// CreateObject创建特定类型的Object结构
func CreateObject(t int, ptr interface{}) (o *GodisObject) {
	o = new(GodisObject)
	o.ObjectType = t
	o.Ptr = ptr
	return
}
