package main

import (
	"fmt"
)

type pc struct {
	ram   int
	brand string
}

func (pc pc) String() string {
	return fmt.Sprintf("La memoria es %d y su marca es %s", pc.ram, pc.brand)
}

func main() {
	nuevaPc := pc{
		ram:   16,
		brand: "Lenovo",
	}
	fmt.Println(nuevaPc)
}
