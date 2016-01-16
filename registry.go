package fbp

import (
	"fmt"
)

type Registry map[string]ComponentConstructor

var (
	registry Registry
)

func init() {
	registry = make(map[string]ComponentConstructor)
}

func Register(descriptor string, c ComponentConstructor) {
	registry[descriptor] = c
}

func Construct(descriptor, id string) (Component, error) {
	constructor, ok := registry[descriptor]
	if ok {
		return constructor(fmt.Sprintf("%s.%s", descriptor, id)), nil
	} else {
		return nil, fmt.Errorf("Component '%s' not registered.", descriptor)
	}
}