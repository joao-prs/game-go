package main

import (
	"os"
    "fmt"
	"time"
	"bufio"
	"strings"
	"io/ioutil"
	"encoding/json"
)


// estruturas de json
type Interface struct {
	Menu1Escolha []string `json:"menu_1_escolha"`
	//HelpMessage    []string `json:"helpMessage"`
	//ExitMessage    []string `json:"exitMessage"`
}


// lê e desserializa o arquivo JSON
func loadTexts(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		return err
	}
	return nil
}
// junta as linhas em um único bloco de texto usando strings.Join
func joinLines(lines []string) string {
	return strings.Join(lines, "\n")
}




func menu_tela_inicial() {

	for {
		clear_screen()
		linha_1()
		fmt.Println("Escolha uma das 3 opções:")
		fmt.Println("	1 - começar o jogo.")
		fmt.Println("	2 - ler o manual do jogo.")
		fmt.Println("	3 - ler sobre.")
		fmt.Println("	q - sair do jogo.")
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
			jogador := panel_create_pessoa(*Pessoa)
			panel_introducao_0(jogador)
			janela_gameplay_0()

		case "2":
			fmt.Println("\nVocê escolheu ler o manual do jogo.")
			time.Sleep(1 * time.Second)
			panel_manual_do_game()
			continue
		
		case "3":
			fmt.Println("\nVocê escolheu ler o manual do jogo.")
			time.Sleep(1 * time.Second)
			panel_sobre_0()
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