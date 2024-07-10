// go mod init game2

/*
	TESTE:
	go run ./main.go ./funcoes_1.go ./panel_guia.go ./panel_menu.go    

	COMPILAR:
	go build -o game2 ./main.go ./funcoes_1.go ./panel_guia.go ./panel_menu.go 
*/

package main

import (
	//"fmt"
	//"os"
	//"bufio"
	//"strconv"
)

func main() {

	menu_tela_inicial()

	// encerrar quando nao houver mais funcoes
	quitando()
}