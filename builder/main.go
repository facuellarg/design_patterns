package main

import (
	"fmt"

	"github.com/facuellarg/desing_patterns/builder/builder"
)

func main() {
	normalBuilder := builder.NewBuilderForceCar()
	director := Director{}
	director.SetBuilder(normalBuilder)
	car := director.Make("WC")
	fmt.Println(&car)
}

type Director struct {
	builder builder.BuilderInterface
}

func (d *Director) SetBuilder(builder builder.BuilderInterface) {
	d.builder = builder
}
func (d *Director) Make(kind string) builder.Car {
	d.builder.BuildMotor()
	if kind == "WC" {
		d.builder.BuildChair()
	}
	return d.builder.Build()

}
