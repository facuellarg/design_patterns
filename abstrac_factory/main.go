package main

import (
	"flag"
	"fmt"
	"log"

	abstracfactory "github.com/facuellarg/desing_patterns/abstrac_factory/abstrac_factory"
)

const (
	modernFlag    = "modern"
	victorianFlag = "victorian"
)

var (
	factoryFlag = flag.String(
		"f",
		modernFlag,
		fmt.Sprintf("%s for modern factory %s for victorian factory", modernFlag, victorianFlag))
)

func main() {
	var factory abstracfactory.AbstractFactory
	flag.Parse()
	switch *factoryFlag {
	case modernFlag:
		factory = abstracfactory.NewModernFactory()
	case victorianFlag:
		factory = abstracfactory.NewVictorianFactory()
	default:
		log.Fatal("invalid flag")
	}

	fmt.Printf("%s\n", factory.CreateChair())
	fmt.Printf("%s\n", factory.CreateSofa())
}
