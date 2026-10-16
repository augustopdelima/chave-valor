package terminal

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func LerTexto(leitor *bufio.Reader, mensagem string) (string, error) {
	fmt.Print(mensagem)

	texto, err := leitor.ReadString('\n')

	if err != nil && len(texto) == 0 {
		return "", err
	}

	return strings.TrimSpace(texto), nil
}

func LerID(leitor *bufio.Reader) (string, error) {
	id, err := LerTexto(leitor, "ID da pessoa:")
	if err != nil {
		return "", err
	}

	if _, err := strconv.ParseUint(id, 10, 64); err != nil {
		return "", fmt.Errorf("Id inválido")
	}

	return id, nil
}
