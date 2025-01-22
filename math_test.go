package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(20, 20)

	if total != 40 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 40)
	}
}
