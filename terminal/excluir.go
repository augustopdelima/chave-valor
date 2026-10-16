package terminal

import (
	"bufio"
	"chave-valor/pessoa"
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ExcluirPessoaPeloTerminal(contxt context.Context, redisDB *redis.Client, leitor *bufio.Reader) error {
	id, err := LerID(leitor)
	if err != nil {
		return err
	}

	confirmacao, err := LerTexto(
		leitor,
		"Tem certeza que deseja excluir? (s/n):",
	)

	if err != nil {
		return err
	}

	if strings.ToLower(confirmacao) != "s" {
		fmt.Println("Exclusão cancelada.")
		return nil
	}

	if err := pessoa.ExcluirPessoa(contxt, redisDB, id); err != nil {
		return err
	}

	fmt.Println("Pessoa excluída com sucesso!")
	return nil
}
