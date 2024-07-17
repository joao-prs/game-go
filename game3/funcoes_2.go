package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func menu_jogador() {

	for {
		//clear_screen()

		fmt.Println("| Escolha uma das 3 opções: |")
		fmt.Println("|  1 - mapa                 |")
		fmt.Println("|  2 - fechar menu          |")
		fmt.Println("|  q - fechar jogo.         |")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nDigite sua ação: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {

		case "1":
			fmt.Printf("\nvoce abriu o mapa\n")
			time.Sleep(1 * time.Second)
			continue

		case "2":
			break

		case "q":
			quitando()

		default:
			fmt.Println("\n\033[91mEscolha inválida!\033[0m")
			time.Sleep(1 * time.Second)
			continue // Recarrega o menu
		}

		// Se chegou até aqui, a escolha foi válida, então sai do loop
		break
	}
}
