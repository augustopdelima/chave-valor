package terminal

import "fmt"

func ExibirMenu() {
	fmt.Println("\n===== GERENCIAMENTO DE PESSOAS =====")
	fmt.Println("1. Cadastrar pessoa")
	fmt.Println("2. Buscar pessoa por ID")
	fmt.Println("3. Listar pessoas")
	fmt.Println("4. Atualizar pessoa")
	fmt.Println("5. Excluir pessoa")
	fmt.Println("0. Sair")
	fmt.Println("====================================")
	fmt.Print("Escolha uma opção: ")
}
