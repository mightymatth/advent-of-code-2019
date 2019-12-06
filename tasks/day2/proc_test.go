package day2

import (
	"testing"
)

func TestProcessor_Start(t *testing.T) {
	p := Processor{}
	p.LoadMemory("input.txt")
	p.SetNounVerb(12, 2)

	res := p.Start()
	
	expectedRes := ProcessorOutput{
		Result: 4690667,
		Noun:   12,
		Verb:   2,
	}

	if res.Result != expectedRes.Result {
		t.Errorf("Expected result to be %v, but got %v\n", expectedRes.Result, res.Result)
	}
}
