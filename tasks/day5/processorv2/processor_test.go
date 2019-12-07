package processorv2

import "testing"

func TestInstruction_Execute(t *testing.T) {
	// equals
	memory := []int{8, 3, 5, 9, 4, 9, 25, 34, 12, 14, 14}
	in := NewInstruction(0, memory)
	in.Execute()

	t.Logf("%v", memory)

	if memory[9] != 1 {
		t.Fail()
	}

	//
	memory = []int{1005, 5, 6, 107, 8, 21, 20, 1006, 20, 31}
	in = NewInstruction(0, memory)
	pointer, _ := in.Execute()

	if pointer != 6 {
		t.Fail()
	}

}

func TestProcessor_Start(t *testing.T) {
	p := Processor{memory: []int{3,21,1008,21,8,20,1005,20,22,107,8,
		21,20,1006,20,31,1106,0,36,98,0,
		0,1002,21,125,20,4,20,1105,1,46,104,
		999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}}

	p.Start()
}
