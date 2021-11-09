package factory

import "fmt"

type Ship struct {
	Name string
}

func (s *Ship) Deliver() error {
	fmt.Println("I'am a deliver ship", s.Name)
	return nil
}
