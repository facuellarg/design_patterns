package factory

import (
	"errors"
	"reflect"
	"sync"
)

type logisticFactory struct {
	mutex   sync.Mutex
	creator map[string]func() Deliverer
}

func NewLogisticFactory() logisticFactory {
	return logisticFactory{
		creator: map[string]func() Deliverer{
			SHIP:  func() Deliverer { return &Ship{} },
			TRUCK: func() Deliverer { return &Truck{} },
		},
	}
}

const (
	SHIP  = "ship"
	TRUCK = "truck"
)

// creator
type LogisticInferface interface {
	CreateTransport(deliver string) (Deliverer, error)
	AddToFactory(deliverer Deliverer, key string) error
}

type Deliverer interface {
	Deliver() error
}

func (lf *logisticFactory) AddToFactory(deliverer Deliverer, key string) error {
	lf.mutex.Lock()
	defer lf.mutex.Unlock()
	if _, ok := lf.creator[key]; ok {
		return errors.New("key is already used")
	}
	delivererType := reflect.TypeOf(deliverer)
	lf.creator[key] = func() Deliverer {
		deliverer := reflect.New(delivererType).Interface().(Deliverer)
		return deliverer
	}
	return nil
}

func (lf *logisticFactory) CreateTransport(deliver string) (Deliverer, error) {
	if creatorFunc, ok := lf.creator[deliver]; !ok {
		return nil, errors.New("invalid input for factor")
	} else {
		return creatorFunc(), nil
	}
}
