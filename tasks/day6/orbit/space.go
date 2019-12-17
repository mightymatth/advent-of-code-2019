package orbit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Space struct {
	com       *Object
	objectMap map[string]*Object
}

func NewSpace(filePath string) Space {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	objectMap := make(map[string]*Object)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		objectNames := strings.Split(strings.TrimSpace(scanner.Text()), ")")

		objectP := getOrCreateObject(objectMap, objectNames[0])
		objectInOrbitP := getOrCreateObject(objectMap, objectNames[1])

		objectInOrbitP.orbits(objectP)
	}

	comObjectP := getOrCreateObject(objectMap, "COM")

	return Space{
		com:       comObjectP,
		objectMap: objectMap,
	}
}

func (s Space) CountOrbitSum() int {
	return s.com.CountOrbitSum()
}

func (s Space) CalculateOrbitalChanges(fromObjectName string, toObjectName string) int {
	changeCount := 0
	from := *s.findObjectByName(fromObjectName).OrbitsIn
	to := *s.findObjectByName(toObjectName).OrbitsIn

	fromDepth := s.getDepth(from)
	toDepth := s.getDepth(to)

	if fromDepth < toDepth {
		for i := 0; i < toDepth-fromDepth; i++ {
			to = *to.OrbitsIn
			changeCount++
		}
	} else {
		for i := 0; i < fromDepth-toDepth; i++ {
			from = *from.OrbitsIn
			changeCount++
		}
	}

	for {
		if from.Name == to.Name || from.OrbitsIn == nil || to.OrbitsIn == nil {
			break
		}
		to = *to.OrbitsIn
		from = *from.OrbitsIn
		changeCount += 2
	}

	return changeCount
}

func (s Space) findObjectByName(name string) Object {
	object, exists := s.objectMap[name]

	if !exists {
		panic(fmt.Sprintf("Cannot find object by name '%v'\n", name))
	}

	return *object
}

func (s Space) getDepth(object Object) int {
	depth := 0

	for {
		if object.Name == "COM" {
			break
		}

		object = *object.OrbitsIn
		depth++
	}

	return depth
}
