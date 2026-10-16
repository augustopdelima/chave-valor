package main

import (
	"bufio"
	"chave-valor/terminal"
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func GerarProximoId(contx context.Context, redisDB *redis.Client) (int64, error) {
	return redisDB.Incr(contx, "id").Result()
}

func main() {
	contx := context.Background()

	redisDb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	leitor := bufio.NewReader(os.Stdin)

	for {
		terminal.ExibirMenu()

		var opcao int
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Opção inválida!")
		}

		switch opcao {

		case 1:
			proximoId, err := GerarProximoId(contx, redisDb)
			if err != nil {
				fmt.Println("Erro ao buscar id:", err)
			}

			idSalvo, err := terminal.CadastrarPessoaPeloTerminal(contx, redisDb, leitor, proximoId)

			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Pessoa cadastrada! ID: %d\n", idSalvo)
			}

		case 2:
			if err := terminal.BuscarPessoaPeloTerminal(contx, redisDb, leitor); err != nil {
				fmt.Println("Erro:", err)
			}

		case 3:
			if err := terminal.ListarPessoasPeloTerminal(contx, redisDb); err != nil {
				fmt.Println("Erro:", err)
			}

		case 4:
			if err := terminal.AtualizarPessoaPeloTerminal(contx, redisDb, leitor); err != nil {
				fmt.Println("Erro:", err)
			}

		case 5:
			if err := terminal.ExcluirPessoaPeloTerminal(contx, redisDb, leitor); err != nil {
				fmt.Println("Erro:", err)
			}

		case 0:
			fmt.Println("Encerrando o programa...")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")

		}

	}

}
