package pessoa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func BuscarPessoa(redisDB *redis.Client, id string, contx context.Context) (Pessoa, error) {

	key := GerarKeyPessoa(id)

	var p Pessoa

	pessoaJson, err := redisDB.Get(contx, key).Bytes()

	if err != nil {

		if errors.Is(err, redis.Nil) {
			return p, fmt.Errorf("pessoa com ID %s não encontrada", id)
		}

		return p, fmt.Errorf("buscar pessoa: %s", id)
	}

	if err := json.Unmarshal(pessoaJson, &p); err != nil {
		return p, fmt.Errorf("desserializar pessoa: %w", err)
	}

	if p.ID == "" {
		p.ID = id
	}

	return p, nil
}

func BuscarPessoas(contx context.Context, redisDB *redis.Client) ([]Pessoa, error) {

	var pessoas []Pessoa

	iter := redisDB.Scan(contx, 0, "pessoa:*", 100).Iterator()

	for iter.Next(contx) {
		key := iter.Val()

		valor, err := redisDB.Get(contx, key).Bytes()

		if err != nil {
			return nil, fmt.Errorf("buscar registro: %w", err)
		}

		var p Pessoa
		if err := json.Unmarshal(valor, &p); err != nil {
			return nil, fmt.Errorf("desserializar pessoa: %w", err)
		}

		if p.ID == "" {
			p.ID = ExtrairIdDaKey(key)
		}

		pessoas = append(pessoas, p)
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("listar pessoas: %w", err)
	}

	sort.Slice(pessoas, func(i, j int) bool {
		idI, errI := strconv.ParseInt(pessoas[i].ID, 10, 64)
		idJ, errJ := strconv.ParseInt(pessoas[j].ID, 10, 64)

		if errI != nil || errJ != nil {
			return pessoas[i].ID < pessoas[j].ID
		}

		return idI < idJ
	})

	return pessoas, nil

}
