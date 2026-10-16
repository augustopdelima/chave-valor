package terminal

import (
	"bufio"
	"chave-valor/pessoa"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func AtualizarPessoaPeloTerminal(context context.Context, redisDB *redis.Client, leitor *bufio.Reader) error {
	id, err := LerID(leitor)
	if err != nil {
		return err
	}

	p, err := ReceberDadosPessoa(leitor)
	if err != nil {
		return err
	}

	if err := pessoa.AtualizarPessoa(context, redisDB, id, p); err != nil {
		return err
	}

	fmt.Printf("Pessoa %s atualizada com sucesso!\n", id)
	return nil
}
