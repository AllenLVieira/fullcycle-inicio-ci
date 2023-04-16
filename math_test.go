package main

import (
	"testing"
	"testing/quick"
)

func TestSomar(t *testing.T) {
	funcaoSomar := func(a, b int) int {
		return Somar(a, b)
	}

	err := quick.CheckEqual(funcaoSomar, func(a, b int) int {
		return a + b
	}, nil)

	if err != nil {
		t.Error(err)
	}
}
