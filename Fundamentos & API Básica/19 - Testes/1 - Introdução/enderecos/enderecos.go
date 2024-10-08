package enderecos

import "strings"

//TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraEndereco := strings.Split(endereco, " ")[0]
	primeiraPalavraEndereco = strings.ToLower(primeiraPalavraEndereco)

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraEndereco
	}
	return "Tipo Inválido"
}
