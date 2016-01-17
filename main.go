package fbp

import (
    "fmt"
)

func main() {
    network := NewNetwork("description")
    err := network.Connect("out", "in")
    if err != nil {
        fmt.Println("Nah... da ginge nie goe eh!")
    }
}