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
		fmt.Println("|  1 - Inventario           |")
		fmt.Println("|  2 - mapa                 |")
		fmt.Println("|  3 - fechar menu          |")
		fmt.Println("|  4 - personagem           |")
		fmt.Println("|  q - sair do jogo.        |\n")


		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite sua ação: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			fmt.Printf("\nvoce abriu o inventario\n")
			time.Sleep(1 * time.Second)
			continue

		case "2":
			fmt.Printf("\nvoce abriu o mapa\n")
			time.Sleep(1 * time.Second)
			continue
		
		case "3":
			//time.Sleep(1 * time.Second)
			break

		
		case "4":
			//time.Sleep(1 * time.Second)
			jogador := panel_create_player()
			menu_jogador(jogador)
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



func menu_jogador(p Pessoa) {

	for {
		clear_screen()

		fmt.Println("Nome: ", p.Nome)
		fmt.Println("vida: ", p.Vida)
		linha_1()
		fmt.Println("| opções:             |")
		fmt.Println("|  1 - fechar menu    |")
		fmt.Println("|  q - sair do jogo.  |\n")


		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite sua ação: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			fmt.Printf("\nvoce abriu o inventario\n")
			time.Sleep(1 * time.Second)
			continue

		case "2":
			fmt.Printf("\nvoce abriu o mapa\n")
			time.Sleep(1 * time.Second)
			continue
		
		case "3":
			//time.Sleep(1 * time.Second)
			break

		
		case "2":
			fmt.Printf("\nvoce abriu o mapa\n")
			time.Sleep(1 * time.Second)
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