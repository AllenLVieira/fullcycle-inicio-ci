package main

import "testing"

func testarSoma(t *testing.T) {
	total := Somar(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma incorreto. O esperado era %d", 30)
	}
}
