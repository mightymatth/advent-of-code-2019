package picture

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPicture(t *testing.T) {
	p := LoadPicture("../inputDummy.txt", 3, 2)

	assert.Equal(t, 3, p.Width)
	assert.Equal(t, 2, p.Height)
	assert.Equal(t, 2, len(p.Layers))
	assert.Equal(t, []int{1, 2, 3}, p.Layers[0].pixels[0])
	assert.Equal(t, []int{4, 5, 6}, p.Layers[0].pixels[1])
	assert.Equal(t, []int{7, 8, 9}, p.Layers[1].pixels[0])
	assert.Equal(t, []int{0, 1, 2}, p.Layers[1].pixels[1])
}

func TestPicture_IntegrityCheck(t *testing.T) {
	p := LoadPicture("../inputDummy.txt", 3, 2)
	assert.Equal(t, 1, p.IntegrityCheck())

	p2 := LoadPicture("../input.txt", 25, 6)
	assert.Equal(t, 2193, p2.IntegrityCheck())
}
