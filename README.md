# Chave-Valor

Aplicação de terminal (CLI) em Go para gerenciamento de pessoas, usando o **Redis** como armazenamento chave-valor.

Cada pessoa é gravada na chave `pessoa:<id>`, e o valor é um JSON com os dados da pessoa — incluindo o próprio `id`.

## Funcionalidades

- Cadastrar pessoa
- Buscar pessoa por ID
- Listar todas as pessoas
- Atualizar pessoa
- Excluir pessoa

## Requisitos

- [Go](https://go.dev/) 1.27 ou superior
- [Redis](https://redis.io/) 7 (há um `Dockerfile` no projeto para subir o servidor)
- (Opcional) [Docker](https://www.docker.com/)

## Como rodar

### 1. Subir o Redis

Com Docker, usando o `Dockerfile` do projeto:

```bash
docker build -t chave-valor-redis .
docker run -d --name chave-valor-redis -p 6379:6379 chave-valor-redis
```

O container roda o Redis com persistência habilitada (`--appendonly yes`).

Se você já tem um Redis local rodando em `localhost:6379`, pode pular esta etapa.

### 2. Rodar a aplicação

```bash
go run .
```

A conexão com o Redis é configurada em `main.go` (endereço `localhost:6379`, banco `0`, sem senha). Ajuste ali caso seu Redis use outra configuração.

## Uso

Ao iniciar, o programa exibe o menu:

```
===== GERENCIAMENTO DE PESSOAS =====
1. Cadastrar pessoa
2. Buscar pessoa por ID
3. Listar pessoas
4. Atualizar pessoa
5. Excluir pessoa
0. Sair
====================================
Escolha uma opção:
```

No cadastro e na atualização são solicitados:

- Nome
- Data de nascimento (`AAAA-MM-DD`)
- E-mail

Todos os campos são obrigatórios.

## Modelo de dados

```go
type Pessoa struct {
	ID             string `json:"id"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"dataNascimento"`
	Email          string `json:"email"`
}
```

Exemplo do valor armazenado no Redis:

```json
{
  "id": "1",
  "nome": "Maria Silva",
  "dataNascimento": "1990-05-20",
  "email": "maria@exemplo.com"
}
```

### Chaves e geração de ID

- Chave: `pessoa:<id>` (ex.: `pessoa:1`).
- Os IDs são gerados com o comando `INCR` na chave `id` do Redis, garantindo valores únicos e incrementais.
- O cadastro usa `SETNX`, evitando sobrescrever um registro existente.
- Registros antigos sem o campo `id` no JSON continuam funcionando: o ID é recuperado a partir da própria chave.

## Estrutura do projeto

```
.
├── main.go              # Ponto de entrada, conexão com o Redis e menu
├── pessoa/              # Regras de negócio e acesso ao Redis
│   ├── pessoa.go        # Struct Pessoa e helpers de chave
│   ├── cadastrar.go     # Cadastro (SETNX)
│   ├── buscar.go        # Busca por ID e listagem (SCAN)
│   ├── atualizar.go     # Atualização (SET)
│   └── excluir.go       # Exclusão (DEL)
├── terminal/            # Interface de linha de comando
│   ├── menu.go          # Exibição do menu
│   ├── leitor.go        # Leitura de texto/ID do terminal
│   ├── cadastrar.go
│   ├── buscar.go
│   ├── atualizar.go
│   └── excluir.go
├── Dockerfile           # Imagem do Redis para desenvolvimento
├── go.mod
└── go.sum
```

## Desenvolvimento

Compilar e verificar o código:

```bash
go build ./...
go vet ./...
```
