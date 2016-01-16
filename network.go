package fbp

import (
	"strings"
)

type Network struct {
	Descriptor string
	Components map[string]Component
}

func NewNetwork(descriptor string) *Network {
	n := new(Network)
	n.Descriptor = descriptor
	n.Components = make(map[string]Component)
	return n
}

func (n *Network) Add(component Component) error {
	n.Components[component.Descriptor()] = component
	return nil
}

func SplitComponentAndPort(s string) (component string, port string){
	e := strings.Split(s, ":")
	component = e[0]
	if len(e) > 1 {
		port = e[1]
	}
	return
}
func (n *Network) Connect(out, in string) error {
	
	return nil
}

