package enderecos

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}
	return "Tipo inválido"
}