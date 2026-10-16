package pessoa

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func AtualizarPessoa(contx context.Context, redisDB *redis.Client, id string, pessoa Pessoa) error {
	key := GerarKeyPessoa(id)

	existe, err := redisDB.Exists(contx, key).Result()

	if err != nil {
		return fmt.Errorf("verificar pessoa: %w", err)
	}

	if existe == 0 {
		return fmt.Errorf("pessoa com ID %s não encontrada", id)
	}

	pessoaJSON, err := json.Marshal(pessoa)
	if err != nil {
		return fmt.Errorf("serializar pessoa: %w", err)
	}

	if err := redisDB.Set(contx, key, pessoaJSON, 0).Err(); err != nil {
		return fmt.Errorf("atualizar pessoa: %w", err)
	}

	return nil

}
