package terminal

import (
	"bufio"
	"chave-valor/pessoa"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ReceberDadosPessoa(leitor *bufio.Reader) (pessoa.Pessoa, error) {

	var p pessoa.Pessoa

	fmt.Print("Nome: ")
	nome, err := leitor.ReadString('\n')
	if err != nil && len(nome) == 0 {
		return pessoa.Pessoa{}, fmt.Errorf("ler nome: %w", err)
	}
	p.Nome = strings.TrimSpace(nome)

	fmt.Print("Data de nascimento (AAAA-MM-DD):")
	dataNascimento, err := leitor.ReadString('\n')
	if err != nil && len(dataNascimento) == 0 {
		return pessoa.Pessoa{}, fmt.Errorf("ler data de nascimento: %w", err)
	}
	p.DataNascimento = strings.TrimSpace(dataNascimento)

	fmt.Print("E-mail:")
	email, err := leitor.ReadString('\n')
	if err != nil && len(email) == 0 {
		return pessoa.Pessoa{}, fmt.Errorf("ler e-mail: %w", err)
	}
	p.Email = strings.TrimSpace(email)

	if p.Nome == "" || p.DataNascimento == "" || p.Email == "" {
		return pessoa.Pessoa{}, fmt.Errorf("todos os campos são obrigatórios")
	}

	return p, nil
}

func CadastrarPessoaPeloTerminal(
	contx context.Context,
	redisDB *redis.Client,
	leitor *bufio.Reader,
	id int64,
) (int64, error) {

	p, err := ReceberDadosPessoa(leitor)
	if err != nil {
		return 0, err
	}

	err = pessoa.CadastrarPessoa(redisDB, p, strconv.FormatInt(id, 10), contx)

	if err != nil {
		return 0, err
	}

	return id, nil

}
