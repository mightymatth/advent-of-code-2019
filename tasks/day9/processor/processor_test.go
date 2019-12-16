package processor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessor(t *testing.T) {
	p := NewProcessor()
	p.LineToMemory("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	p.AppendFreeMemory(120)

	p.Start(nil)

	outputs := make([]int, 0)
	for output := range p.Output {
		outputs = append(outputs, output)
	}

	assert.Equal(t, 109, outputs[0])
	assert.Equal(t, 1, outputs[1])
	assert.Equal(t, 0, outputs[len(outputs)-2])
	assert.Equal(t, 99, outputs[len(outputs)-1])
}

func TestProcessor2(t *testing.T) {
	p := NewProcessor()
	p.LineToMemory("1102,34915192,34915192,7,4,7,99,0")
	p.AppendFreeMemory(120)

	p.Start(nil)

	outputs := make([]int, 0)
	for output := range p.Output {
		outputs = append(outputs, output)
	}

	assert.Equal(t, 1219070632396864, outputs[0])
}

func TestProcessor3(t *testing.T) {
	p := NewProcessor()
	p.LineToMemory("104,1125899906842624,99")
	p.AppendFreeMemory(120)

	p.Start(nil)

	outputs := make([]int, 0)
	for output := range p.Output {
		outputs = append(outputs, output)
	}

	assert.Equal(t, 1125899906842624, outputs[0])
}

func TestProcessorRealInputPart1(t *testing.T) {
	p := NewProcessor()
	p.LoadMemory("../input.txt")
	p.AppendFreeMemory(100)

	p.Input <- 1
	p.Start(nil)

	outputs := make([]int, 0)
	for output := range p.Output {
		outputs = append(outputs, output)
	}

	assert.Equal(t, 3454977209, outputs[0])
}

func TestProcessorRealInputPart2(t *testing.T) {
	p := NewProcessor()
	p.LoadMemory("../input.txt")
	p.AppendFreeMemory(2000)

	p.Input <- 2
	p.Start(nil)

	outputs := make([]int, 0)
	for output := range p.Output {
		outputs = append(outputs, output)
		fmt.Println(output)
	}

	assert.Equal(t, 50120, outputs[0])
}