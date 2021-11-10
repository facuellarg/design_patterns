package abstracfactory

type modernChair struct {
	legs   int
	canSit bool
}

func NewModernChair() modernChair {
	return modernChair{
		4,
		true,
	}
}

func (mc modernChair) Legs() int {
	return mc.legs
}
func (mc modernChair) SitOn() bool {
	defer func() { mc.canSit = false }()
	return mc.canSit
}
func (mc modernChair) CanSitOn() bool {
	return mc.canSit
}

func (_ modernChair) String() string {
	return "i am a modern chair"
}
