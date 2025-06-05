//TESTE DE UNIDADE

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Rodovia dos Imigrantes", "rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada Qualquer", "estrada"},
		{"RUA DOS BOBOS", "rua"},
		{"AVENIDA REBOUCAS", "avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.enderecoEsperado)
		}

	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
