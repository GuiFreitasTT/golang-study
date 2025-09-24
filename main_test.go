package main

import "testing"

func TestSomar(t *testing.T) {
    resultado := Somar(10, 10)
    esperado := 20

    if resultado != esperado {
        t.Errorf("Esperado %d, mas obteve %d", esperado, resultado)
    }
}
