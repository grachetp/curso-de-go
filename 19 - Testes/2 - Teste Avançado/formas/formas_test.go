package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			//t.Errorf("A àrea recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) //Teste vai dar erro, mas vai continuar a executar os restantes.
			t.Fatalf("A àrea recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) //Teste vai dar erro e vai parar aqui todos os testes sendo executados.
		}
	})

	t.Run("Circúlo", func(t *testing.T) {
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			//t.Errorf("A àrea recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
			t.Fatalf("A àrea recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) //Teste vai dar erro e vai parar aqui todos os testes sendo executados.
		}
	})
}
