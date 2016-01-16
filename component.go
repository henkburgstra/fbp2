package fbp


type ComponentConstructor func (descriptor string) Component

type Component interface {
	Descriptor() string
}