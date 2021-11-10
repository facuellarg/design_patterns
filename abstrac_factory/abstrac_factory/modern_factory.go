package abstracfactory

type modernFactory struct{}

func NewModernFactory() modernFactory {
	return modernFactory{}
}

func (vf modernFactory) CreateChair() Chair {
	return NewModernChair()
}

func (vf modernFactory) CreateSofa() Sofa {
	return NewModernSofa()
}
