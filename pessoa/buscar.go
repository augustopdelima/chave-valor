package pessoa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func BuscarPessoa(redisDB *redis.Client, id string, contx context.Context) (Pessoa, error) {

	key := GerarKeyPessoa(id)

	var pessoa Pessoa

	pessoaJson, err := redisDB.Get(contx, key).Bytes()

	if err != nil {

		if errors.Is(err, redis.Nil) {
			return pessoa, fmt.Errorf("pessoa com ID %s não encontrada", id)
		}

		return pessoa, fmt.Errorf("buscar pessoa: %s", id)
	}

	if err := json.Unmarshal([]byte(pessoaJson), &pessoa); err != nil {
		return pessoa, fmt.Errorf("desserializar pessoa: %w", err)
	}

	return pessoa, nil
}

func BuscarPessoas(contx context.Context, redisDB *redis.Client) ([]Pessoa, error) {

	var pessoas []Pessoa

	iter := redisDB.Scan(contx, 0, "pessoa:*", 100).Iterator()

	for iter.Next(contx) {
		valor, err := redisDB.Get(contx, iter.Val()).Bytes()

		if err != nil {
			return nil, fmt.Errorf("buscar registro: %w", err)
		}

		var pessoa Pessoa
		if err := json.Unmarshal(valor, &pessoa); err != nil {
			return nil, fmt.Errorf("desserializar pessoa: %w", err)
		}

		pessoas = append(pessoas, pessoa)
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("listar pessoas: %w", err)
	}

	return pessoas, nil

}
