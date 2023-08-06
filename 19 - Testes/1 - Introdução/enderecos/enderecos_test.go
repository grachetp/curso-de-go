// Teste de Unidade
// Teste Unitário

package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// TestNomeDaFuncaoATestar(ponteiro do tipo T, pacote padrão para desenvolver testes unitários)
	enderecoParaTeste := "Avenida Tãpinuá"
	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		//t.Error("O tipo recebido é diferente do esperado!")
		t.Errorf("O tipo recebido é diferente do esperado!. Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}

//commando> go test na pasta do teste
