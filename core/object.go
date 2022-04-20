package core

// GodisObject是对特定的数据类型的包装
type GodisObject struct {
	ObjectType int
	Ptr        interface{}
}

const ObjectTypeString = 1

//CreateObject 创建特定类型的object结构
func CreateObject(t int, ptr interface{}) (o *GodisObject) {
	o = new(GodisObject)
	o.ObjectType = t
	o.Ptr = ptr
	return
}
