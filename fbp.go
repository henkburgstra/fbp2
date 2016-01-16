package fbp


type Process struct {
	descriptor string
	ports map[string]*Port
}

func NewProcess(descriptor string) Component {
	p := new(Process)
	p.descriptor = descriptor
	p.ports = make(map[string]*Port)
	return p
}

func (p *Process) Descriptor() string {
	return p.descriptor
}

func (p *Process) AddPort(*Port) {

}
