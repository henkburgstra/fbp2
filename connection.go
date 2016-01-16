package fbp


type Connection struct {
	Capacity int
	From *Port
	To *Port
	Pipe chan interface{}
}

func NewConnection(from, to *Port) *Connection {
	c := new(Connection)
	c.From = from
	c.To = to
	c.Capacity = 1
	c.Pipe = make(chan interface{}, c.Capacity)
	return c
}

func (c *Connection) SetCapacity(cap int) error {
	c.Capacity = cap
	c.Pipe = make(chan interface{}, cap)
	return nil
}

