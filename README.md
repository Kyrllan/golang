# Plano de Estudos Go

## 1. Módulos em Go

### 1.1 Criação e Gerenciamento de Módulos

- **Inicialização de Módulo**:

```bash
go mod init nome-do-modulo
```

### 1.2 Estrutura de Módulos

- Organização de diretórios  
- Subpacotes  
- Arquivos principais e auxiliares  

### 1.3 Exportação e Importação de Funções

- **Funções Exportadas** (Primeira letra maiúscula):

```go
func Escrever() {
    // código aqui
}
```

- **Funções Internas** (Primeira letra minúscula):

```go
func escrever() {
    // código aqui
}
```

## 2. Gerenciamento de Dependências

### 2.1 Instalação de Pacotes

```bash
go get github.com/nome-do-pacote
```

### 2.2 Uso de Pacotes de Terceiros

```go
import (
    "github.com/nome-do-pacote"
)
```

### 2.3 Manutenção de Dependências

- **Remover dependências não utilizadas**:

```bash
go mod tidy
```

## 3. Comandos Essenciais

### 3.1 Gerenciamento de Módulos

```bash
# Iniciar novo módulo
go mod init

# Limpar dependências não utilizadas
go mod tidy

# Criar pasta vendor com dependências
go mod vendor
```

### 3.2 Build e Execução

```bash
# Gerar executável
go build

# Executar projeto
go run .
```
