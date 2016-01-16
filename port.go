package fbp

type PortStatus int

const (
	Open PortStatus = iota
	Closed
)

type PortType int

const (
	In PortType = iota
	Out
)

type Port struct {
	Descriptor  string
	PortType    PortType
	Connections []*Connection
	Status      PortStatus
}

func NewPort(descriptor string, portType PortType) *Port {
	p := new(Port)
	p.PortType = portType
	p.Connections = make([]*Connection, 0, 0)
	return p
}

func (p *Port) Connect(c *Connection) {
	p.Connections = append(p.Connections, c)
} 