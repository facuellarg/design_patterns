package abstracfactory

type victorianSofa struct {
	spots    int
	hasSpots bool
}

func NewVictorianSofa() victorianSofa {
	return victorianSofa{
		3,
		true,
	}
}
func (ms victorianSofa) Spots() int {
	return ms.spots
}

func (ms victorianSofa) HasSpots() bool {
	return ms.hasSpots
}
func (ms victorianSofa) Sit() bool {
	if !ms.HasSpots() {
		return false
	}
	ms.spots--
	ms.hasSpots = ms.spots > 0
	return true
}

func (_ victorianSofa) String() string {
	return "i am a victorian sofa"
}
