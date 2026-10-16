package terminal

import (
	"bufio"
	"chave-valor/pessoa"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func BuscarPessoaPeloTerminal(
	contxt context.Context,
	redisDB *redis.Client,
	leitor *bufio.Reader,
) error {

	id, err := LerID(leitor)
	if err != nil {
		return err
	}

	p, err := pessoa.BuscarPessoa(redisDB, id, contxt)
	if err != nil {
		return err
	}

	fmt.Println("\n--- Pessoa encontrada ---")
	fmt.Println("ID:", id)
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Data de nascimento:", p.DataNascimento)
	fmt.Println("E-mail:", p.Email)

	return nil

}

func ListarPessoasPeloTerminal(contxt context.Context, redisDB *redis.Client) error {
	pessoas, err := pessoa.BuscarPessoas(contxt, redisDB)
	if err != nil {
		return err
	}

	if len(pessoas) == 0 {
		fmt.Println("Nenhuma pessoa cadastrada.")
		return nil
	}

	for i, pessoa := range pessoas {
		fmt.Printf("\n--- Pessoa %d ---\n", i+1)
		fmt.Println("Nome:", pessoa.Nome)
		fmt.Println("Data de nascimento:", pessoa.DataNascimento)
		fmt.Println("E-mail:", pessoa.Email)
	}

	return nil
}
