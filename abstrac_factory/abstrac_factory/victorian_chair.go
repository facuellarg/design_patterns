package abstracfactory

type victorianChair struct {
	legs   int
	canSit bool
}

func NewVictorianChair() victorianChair {
	return victorianChair{
		3,
		true,
	}
}

func (vc victorianChair) Legs() int {
	return vc.legs
}
func (vc victorianChair) SitOn() bool {
	defer func() { vc.canSit = false }()
	return vc.canSit
}
func (vc victorianChair) CanSitOn() bool {
	return vc.canSit
}

func (_ victorianChair) String() string {
	return "i am a victorian chair"
}
