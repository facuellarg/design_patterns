package abstracfactory

type Chair interface {
	CanSitOn() bool
	Legs() int
	SitOn() bool
}
