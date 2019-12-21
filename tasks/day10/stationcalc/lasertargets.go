package stationcalc

type LaserTargets struct {
	Laser   *Asteroid
	Targets []Asteroid
}

func (lt LaserTargets) Len() int {
	return len(lt.Targets)
}

func (lt LaserTargets) Less(i, j int) bool {
	laserPosition := lt.Laser.Position

	angle1 := measureAngle(laserPosition, lt.Targets[i].Position)
	angle2 := measureAngle(laserPosition, lt.Targets[j].Position)

	return angle1 < angle2
}

func (lt LaserTargets) Swap(i, j int) {
	lt.Targets[i], lt.Targets[j] = lt.Targets[j], lt.Targets[i]
}
