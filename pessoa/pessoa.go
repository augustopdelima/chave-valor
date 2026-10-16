package pessoa

type Pessoa struct {
	Nome           string `json:"nome"`
	DataNascimento string `json:"dataNascimento"`
	Email          string `json:"email"`
}

func GerarKeyPessoa(id string) string {
	return "pessoa:" + id
}
