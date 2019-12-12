package picture

type Layer struct {
	pixels [][]int
}

func NewLayer(width int, height int) Layer {
	pixels := make([][]int, height)
	for i := range pixels {
		pixels[i] = make([]int, width)
	}

	return Layer{pixels: pixels}
}
