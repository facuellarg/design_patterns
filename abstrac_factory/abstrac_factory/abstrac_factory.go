package abstracfactory

type AbstractFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}
