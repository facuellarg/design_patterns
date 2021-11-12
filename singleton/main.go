package main

import (
	"fmt"
	"sync"

	"github.com/facuellarg/desing_patterns/singleton/singleton"
)

func main() {
	count := 1000
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			singleton.GetManager().Add()
			wg.Done()
		}()
	}
	wg.Wait()
	manager2 := singleton.GetManager()
	fmt.Println(manager2.GetValue())
}
