// Teste de Unidade
// Teste Unitário

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// // TestNomeDaFuncaoATestar(ponteiro do tipo T, pacote padrão para desenvolver testes unitários)
	// enderecoParaTeste := "Avenida Tãpinuá"
	// tipoDeEnderecoEsperado := "Avenida"

	// tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)

	// if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
	// 	t.Error("O tipo recebido é diferente do esperado!")
	// 	t.Errorf("O tipo recebido é diferente do esperado!. Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	// }

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida das Rebouças", "avenida"},
		{"Estrada JK", "estrada"},
		{"Rodovia BR-364", "rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Boteco das Rosas", "Tipo Inválido"},
		{"RUA das Rosas", "rua"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado!. Esperava %s e recebeu %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}

}

//commando> go test na pasta do teste
