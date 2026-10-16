package pessoa

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func CadastrarPessoa(redisDb *redis.Client, pessoa Pessoa, id string, contx context.Context) error {

	key := GerarKeyPessoa(id)

	pessoaJson, err := json.Marshal(pessoa)

	if err != nil {
		return fmt.Errorf("serializar pessoa: %w", err)
	}

	cadastrada, err := redisDb.SetNX(contx, key, pessoaJson, 0).Result()

	if err != nil {
		return fmt.Errorf("cadastrar pessoa: %w", err)
	}

	if !cadastrada {
		return fmt.Errorf("pessoa com ID %s já existe", id)
	}

	return nil

}
