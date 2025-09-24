package main

import "testing"

func TestSomar(t *testing.T) {
    resultado := Somar(5, 2)
    esperado := 4

    if resultado != esperado {
        t.Errorf("Esperado %d, mas obteve %d", esperado, resultado)
    }
}
