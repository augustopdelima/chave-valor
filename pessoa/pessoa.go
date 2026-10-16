package pessoa

import "strings"

type Pessoa struct {
	ID             string `json:"id"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"dataNascimento"`
	Email          string `json:"email"`
}

func GerarKeyPessoa(id string) string {
	return "pessoa:" + id
}

func ExtrairIdDaKey(key string) string {
	return strings.TrimPrefix(key, "pessoa:")
}
