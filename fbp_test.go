package fbp

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	descriptor := "test"
	Register(descriptor, NewProcess)
	if _, ok := registry[descriptor]; !ok {
		t.Errorf("Constructor %s is registered but not in the registry.")
	}
}

func TestConstruct(t *testing.T) {
	component := "test"
	id := "test"
	descriptor := fmt.Sprintf("%s.%s", component, id)
	Register(component, NewProcess)
	p, err := Construct(component, id)
	if err != nil {
		t.Errorf("Process %s could not be constructed: %s", descriptor, err.Error())
		return
	}
	if p.Descriptor() != descriptor {
		t.Errorf("Component %s has descriptor %s, %s expected", component, p.Descriptor(), descriptor)
	}	
}

func TestNewNetwork(t *testing.T) {
	nd := "test"
	network := NewNetwork(nd)
	if network == nil {
		t.Error("nil returned, Network instance expected")
		return
	}
	if network.Descriptor != nd {
		t.Errorf("Network has descriptor %s, %s expected", network.Descriptor, nd)
	}
}

func TestNetworkAdd(t *testing.T) {
	nd := "testnetwork"
	component := "test"
	id := "test"
	descriptor := fmt.Sprintf("%s.%s", component, id)
	Register(component, NewProcess)
	p, err := Construct(component, id)
	if err != nil {
		t.Errorf("Process %s could not be constructed: %s", descriptor, err.Error())
		return
	}
	network := NewNetwork(nd)
	err = network.Add(p)
	if err != nil {
		t.Errorf("Could not add process %s to the network: %s", p.Descriptor(), err.Error())
	}
}

func TestSplitComponentAndPort(t *testing.T) {
	component := "test.1"
	port := "in"
	descriptor := fmt.Sprintf("%s:%s", component, port)
	c, p := SplitComponentAndPort(descriptor)
	if c != component {
		t.Errorf("Descriptor %s, expected component %s, got %s", descriptor, component, c)
	}
	if p != port {
		t.Errorf("Descriptor %s, expected port %s, got %s", descriptor, port, p)
	}
	c, p = SplitComponentAndPort(component)
	if c != component {
		t.Errorf("Descriptor %s, expected component %s, got %s", component, component, c)
	}
	if p != "" {
		t.Errorf("Descriptor %s, expected empty port, got %s", component, p)
	}
}