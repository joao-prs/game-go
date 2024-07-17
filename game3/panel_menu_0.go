package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func menu_tela_inicial() {

	for {
		clear_screen()
		linha_1()
		fmt.Println("Escolha uma das 3 opções:")
		fmt.Println("\n   1 - começar o jogo.")
		fmt.Println("   2 - ler o manual do jogo.")
		fmt.Println("   3 - ler sobre.")
		fmt.Println("   q - sair do jogo.")
		linha_1()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite sua escolha: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			fmt.Println("\nVocê escolheu começar o jogo.")
			time.Sleep(1 * time.Second)

			// o jogo inteiro ta aqui kkkkkkkkkkkkkkkkkk vou arrumar depois
			tela_gameplay_0()

		case "2":
			fmt.Println("\nVocê escolheu ler o manual do jogo.")
			time.Sleep(1 * time.Second)
			panel_0_manual_do_game()
			continue

		case "3":
			fmt.Println("\nVocê escolheu ler o manual do jogo.")
			time.Sleep(1 * time.Second)
			panel_0_sobre()
			continue

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
