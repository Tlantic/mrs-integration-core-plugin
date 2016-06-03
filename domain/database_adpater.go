package domain


type DbAdapter interface {
	GetName() string
	Create(DbObject) error
	Read(string) (error, []*DbObject)
	ReadOne(string) (error, *DbObject)
	ReadOneWithType(string, interface{}) (error, *DbObject)
	Update(*DbObject) error
	UpdateOne(*DbObject) error
	Destroy(*DbObject) error
	DestroyOne(string) error
}


type DbObject struct {
	Key      string
	Data     interface{}
	Expiry   uint32
}

func NewDbObject(key string) *DbObject {
	return &DbObject{
		Key: key,
	}
}

func (c *DbObject) GetKey() string {
	return c.Key
}

func (c *DbObject) SetKey(key string) {
	c.Key = key
}

func (c *DbObject) SetData(data interface{}) {
	c.Data = data
}

func (c *DbObject) GetData() interface{} {
	return c.Data
}


func (c *DbObject) SetExpiry(time uint32) {
	c.Expiry = time
}