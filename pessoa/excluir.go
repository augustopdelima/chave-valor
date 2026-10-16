package pessoa

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExcluirPessoa(contx context.Context, redisDB *redis.Client, id string) error {
	key := GerarKeyPessoa(id)

	removidas, err := redisDB.Del(contx, key).Result()

	if err != nil {
		return fmt.Errorf("excluir pessoa: %w ", err)
	}

	if removidas == 0 {
		return fmt.Errorf("pessoa com ID %s não encontrada", id)
	}

	return nil
}
