package abstracfactory

type modernSofa struct {
	spots    int
	hasSpots bool
}

func NewModernSofa() modernSofa {
	return modernSofa{
		2,
		true,
	}
}
func (ms modernSofa) Spots() int {
	return ms.spots
}

func (ms modernSofa) HasSpots() bool {
	return ms.hasSpots
}
func (ms modernSofa) Sit() bool {
	if !ms.HasSpots() {
		return false
	}
	ms.spots--
	ms.hasSpots = ms.spots > 0
	return true
}
func (_ modernSofa) String() string {
	return "i am a modern sofa"
}
