package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func janela_gameplay_0() {

	for {

		fmt.Println("	m -[Menu]   ")
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

		case "":
			//fmt.Printf("hahahhahahahahahahahahahah")
			//time.Sleep(1 * time.Second)
			time.Sleep(100 * time.Millisecond)
			break

		default:
			fmt.Println("\n\033[91mEscolha inválida!\033[0m")
			//time.Sleep(1 * time.Second)
			time.Sleep(100 * time.Millisecond)
			continue // Recarrega o menu
		}
		// Se chegou até aqui, a escolha foi válida, então sai do loop
		break
	}

}

func tela_gameplay_0() {
	clear_screen()
	linha_0()
	dialog_box_0_scene_0()
	confirma_enter()
	//janela_gameplay_0()

	clear_screen()
	linha_0()
	dialog_box_0_scene_1()
	confirma_enter()
	//janela_gameplay_0()

	clear_screen()
	linha_0()
	dialog_box_0_scene_2()
	confirma_enter()
	//janela_gameplay_0()

	clear_screen()
	linha_0()
	dialog_box_0_scene_3()
	confirma_enter()
	//janela_gameplay_0()

	clear_screen()
	linha_0()
	dialog_box_0_scene_4()
	confirma_enter()
	//janela_gameplay_0()
}
