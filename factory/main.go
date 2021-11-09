package main

import (
	"fmt"
	"log"

	"github.com/facuellarg/desing_patterns/factory/factory"
)

type MyDeliverer struct {
	Speed float64
}

func (md MyDeliverer) Deliver() error {
	fmt.Println("I am a new deliverer, mi speed is:", md.Speed)
	return nil
}

func main() {
	lf := factory.NewLogisticFactory()
	if err := lf.AddToFactory(MyDeliverer{}, "myDeliverer"); err != nil {
		log.Fatal(err)
	}
	transport, _ := lf.CreateTransport(factory.SHIP)
	transport2, _ := lf.CreateTransport(factory.TRUCK)
	myDeliverer, _ := lf.CreateTransport("myDeliverer")
	md := myDeliverer.(*MyDeliverer)
	md.Speed = 100
	transport.Deliver()
	transport2.Deliver()
	myDeliverer.Deliver()
}
