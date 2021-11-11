package builder

type Car struct {
	chair       Chair
	antenna     Antenna
	doors       []Door
	frontBumper FrontBumper
	rearBumper  RearBumper
	muffler     Muffler
	spoiler     Spoiler
	wheels      []Wheel
	motor       Motor
}

func (c *Car) SetChair(chair Chair) {
	c.chair = chair
}

func (c *Car) SetAntenna(antenna Antenna) {
	c.antenna = antenna
}

func (c *Car) SetDoors(doors []Door) {
	c.doors = doors
}

func (c *Car) Setwheels(wheels []Wheel) {
	c.wheels = wheels
}
func (c *Car) SetMotor(motor Motor) {
	c.motor = motor
}

func (c *Car) SetSpoiler(spoiler Spoiler) {
	c.spoiler = spoiler
}
func (c *Car) SetMuffler(muffler Muffler) {
	c.muffler = muffler
}

func (c *Car) SetRearBumper(rearBumper RearBumper) {
	c.rearBumper = rearBumper
}

func (c *Car) SetFrontBumper(frontBumper FrontBumper) {
	c.frontBumper = frontBumper
}

func (c *Car) String() string {
	var carString string
	if c.motor.Force > 100 {
		carString = "I am a speed car"
	} else {
		carString = "I am a normal car"
	}
	if c.chair.color == "" {
		return carString + " and I don't have chair color"
	}
	return carString + " and my chair color is " + c.chair.color
}
