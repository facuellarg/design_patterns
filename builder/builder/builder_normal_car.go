package builder

type builderNormalCar struct {
	car Car
}

func NewBuilderNormalCar() *builderNormalCar {
	return &builderNormalCar{Car{}}
}

func (bf *builderNormalCar) Build() Car {
	buildedCar := bf.car
	bf.Reset()
	return buildedCar
}
func (bf *builderNormalCar) Reset() {
	bf.car = Car{}
}
func (bf *builderNormalCar) BuildMotor() {
	bf.car.SetMotor(Motor{90})
}
func (bf *builderNormalCar) BuildChair() {
	bf.car.SetChair(Chair{"blue"})
}
