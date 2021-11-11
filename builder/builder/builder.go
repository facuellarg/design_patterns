package builder

type BuilderInterface interface {
	Reset()
	BuildChair()
	BuildMotor()
	Build() Car
}
