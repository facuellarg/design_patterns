package abstracfactory

type victorianFactory struct{}

func NewVictorianFactory() victorianFactory {
	return victorianFactory{}
}

func (vf victorianFactory) CreateChair() Chair {
	return NewVictorianChair()
}

func (vf victorianFactory) CreateSofa() Sofa {
	return NewVictorianSofa()
}
