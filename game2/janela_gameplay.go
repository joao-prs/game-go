package main

import (
    "fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

func janela_gameplay_0() {

	for {
		clear_screen()
		linha_0()
		fmt.Println("textotextotextotextotexto")
		fmt.Println("textotextotextotextotexto")
		fmt.Println("textotextotextotextotexto\n")

		//fmt.Println("	m -[Menu]  q - [sair do jogo.]")
		fmt.Println("	m -[Menu]")
		linha_0()

		// ler a resposta
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite sua escolha: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "m":
			menu_jogador()
			continue

		//case "q":
		//	quitando()

		default:
			fmt.Println("\n\033[91mEscolha inválida!\033[0m")
			time.Sleep(1 * time.Second) 
			continue // Recarrega o menu
		}
		// Se chegou até aqui, a escolha foi válida, então sai do loop
		break
	}
	
}