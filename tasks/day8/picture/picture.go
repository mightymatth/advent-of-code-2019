package picture

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Picture struct {
	Width  int
	Height int
	Layers []Layer
}

func LoadPicture(filePath string, width int, height int) Picture {
	file, err := os.Open(filePath)

	if err != nil {
		panic("Cannot read the file!")
	}
	defer file.Close()

	pixels := make([]int, 0)
	picture := Picture{Width: width, Height: height}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pixelsStr := strings.Split(strings.TrimSpace(scanner.Text()), "")

		for _, pixelStr := range pixelsStr {
			pixel, err := strconv.Atoi(pixelStr)

			if err != nil {
				panic("Cannot convert text to integer!")
			}

			pixels = append(pixels, pixel)
		}
	}
	picture.setPixels(pixels)

	return picture
}

func (p *Picture) setPixels(pixels []int) {
	layersCount := len(pixels) / (p.Width * p.Height)
	layers := make([]Layer, layersCount)
	pixelIndex := 0

	for layerIndex, layer := range layers {
		layer = NewLayer(p.Width, p.Height)

		for i := 0; i < p.Height; i++ {
			for j := 0; j < p.Width; j++ {
				layer.pixels[i][j] = pixels[pixelIndex]
				pixelIndex++
			}
		}

		layers[layerIndex].pixels = layer.pixels
	}

	p.Layers = layers
}

func (p Picture) IntegrityCheck() int {
	layer := p.minZeroDigitsLayer()

	return layer.countDigits(1) * layer.countDigits(2)
}

func (l Layer) countDigits(digit int) int {
	counter := 0
	for i := 0; i < len(l.pixels); i++ {
		for j := 0; j < len(l.pixels[i]); j++ {
			if l.pixels[i][j] == digit {
				counter++
			}
		}
	}

	return counter
}

func (p Picture) minZeroDigitsLayer() Layer {
	minZeroDigits := math.MaxInt32
	var minZeroDigitsLayer Layer

	for _, layer := range p.Layers {
		if layerZeroDigits := layer.countDigits(0); layerZeroDigits < minZeroDigits {
			minZeroDigits = layerZeroDigits
			minZeroDigitsLayer = layer
		}
	}

	return minZeroDigitsLayer
}
