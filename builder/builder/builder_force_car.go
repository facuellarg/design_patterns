package builder

type builderForceCar struct {
	car Car
}

func NewBuilderForceCar() *builderForceCar {
	return &builderForceCar{Car{}}
}

func (bf *builderForceCar) Build() Car {
	buildedCar := bf.car
	bf.Reset()
	return buildedCar
}
func (bf *builderForceCar) Reset() {
	bf.car = Car{}
}
func (bf *builderForceCar) BuildMotor() {
	bf.car.SetMotor(Motor{200})
}
func (bf *builderForceCar) BuildChair() {
	bf.car.SetChair(Chair{"red"})
}
