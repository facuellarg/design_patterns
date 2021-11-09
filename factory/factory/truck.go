package factory

import "fmt"

type Truck struct {
	Name string
}

func (t *Truck) Deliver() error {
	fmt.Println("I'am a deliver truck", t.Name)
	return nil
}
