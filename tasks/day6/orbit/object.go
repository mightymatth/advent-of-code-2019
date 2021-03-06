package orbit

type Object struct {
	Name     string
	OrbitsIn *Object
	Orbit    []*Object
}

func (o *Object) orbits(object *Object) {
	object.Orbit = append(object.Orbit, o)
	o.OrbitsIn = object
}

func getOrCreateObject(objectMap map[string]*Object, objectName string) *Object {
	objectP, exists := objectMap[objectName]

	if exists {
		return objectP
	}

	newObjectP := &Object{
		Name:     objectName,
		OrbitsIn: nil,
		Orbit:    make([]*Object, 0),
	}

	objectMap[objectName] = newObjectP

	return newObjectP
}

func (o Object) CountOrbitSum() int {
	sum := 0
	if len(o.Orbit) == 0 {
		return o.countOrbits()
	}

	for _, object := range o.Orbit {
		sum += object.CountOrbitSum()
	}

	if o.Name == "COM" {
		return sum
	}

	return o.countOrbits() + sum
}

func (o Object) countOrbits() int {
	directOrbits := 0
	indirectOrbits := 0

	if o.OrbitsIn != nil {
		directOrbits = 1
	}

	for _, object := range o.Orbit {
		indirectOrbits += object.countOrbits()
	}

	return directOrbits + indirectOrbits
}
