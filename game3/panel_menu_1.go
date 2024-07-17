package main

import (
	"fmt"
	"time"
)

func panel_introducao_0() {

	//var p Pessoa
	clear_screen()
	linha_0_efeito()
	//linha_0()
	fmt.Printf("Bem vindo, %s!\n\n")
	fmt.Printf("Criei esse prototipo interativo para fins de estudo e humor,\n")
	fmt.Printf("leia todas as regras com calma, para que o programa seja\n")
	fmt.Printf("executado sem nenhum (ou com poucos) problemas. espero que\n")
	fmt.Printf("se divirta.\n")
	fmt.Printf("                                           - github@joao-prs\n")
	linha_0_efeito()
	//linha_0()

	time_3_sec()
	confirma_enter()
}

func panel_0_manual_do_game() {
	var CONFIRM bool

	for true {
		clear_screen()
		linha_0()

		fmt.Printf("Regras e Avisos do JOGO\n\n")
		fmt.Printf("Nao se preocupe se quando aparecer uma caixa de texto o que\n")
		fmt.Printf("voce digitar nao estiver aparencendo, eh normal, eh um bug\n")
		fmt.Printf("que nao consegui resolver, mas tudo que voce digitar estara\n")
		fmt.Printf("sendo lido normalmente pelo seu shell.  :)\n\n")
		fmt.Printf("Sempre que voce ver textos em \033[033mamarelo\033[0m significa que voce pode\n")
		fmt.Printf("interagir com seu teclado usando teclas sugeridas pelo prompt.\n\n")
		fmt.Printf("Quando voce ver textos em \033[32mverde\033[0m significa que são nome de\n")
		fmt.Printf("personagens, quando voce ver textos em \033[94mazul\033[0m sao itens que\n")
		fmt.Printf("você vai interagir, e quando voce ver \033[96mciano\033[0m eh apenas alguma\n")
		fmt.Printf("informacao importante no game. E por ultimo, quando voce ver\n")
		fmt.Printf("\033[91mvermelho\033[0m, significa que sao coisas como sua vida ou algo tao\n")
		fmt.Printf("relevante quanto.\n\n")
		fmt.Printf("                    Enter (avança)\n")
		fmt.Printf("                    y (concorda)\n")
		fmt.Printf("                    n (discorda)\n")
		fmt.Printf("                    q (quita do game)\n\n")

		linha_0()
		time_3_sec()
		//confirma_enter()
		fmt.Printf("\rVoce entendeu ? ")
		CONFIRM = confirma_y_ou_n(CONFIRM)
		if CONFIRM {
			fmt.Printf("\n[voce] sim\n")
			time.Sleep(1 * time.Second)
			break
		} else {
			fmt.Printf("\n[voce] nao\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("\n\ntudo bem, eu explico\n")
			time.Sleep(1 * time.Second)
		}

	}
	clear_screen()
}

func panel_0_sobre() {
	var CONFIRM bool

	for true {
		clear_screen()
		linha_0()

		fmt.Println("MIT License\n")
		fmt.Println("Copyright (c) 2024 João Pedro Ribeiro\n")
		fmt.Println("Permission is hereby granted, free of charge, to any person obtaining a copy")
		fmt.Println("of this software and associated documentation files (the 'Software'), to deal")
		fmt.Println("in the Software without restriction, including without limitation the rights")
		fmt.Println("to use, copy, modify, merge, publish, distribute, sublicense, and/or sell")
		fmt.Println("copies of the Software, and to permit persons to whom the Software is")
		fmt.Println("furnished to do so, subject to the following conditions:\n")
		fmt.Println("The above copyright notice and this permission notice shall be included in all")
		fmt.Println("copies or substantial portions of the Software.\n")
		fmt.Println("THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR")
		fmt.Println("IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,")
		fmt.Println("FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE")
		fmt.Println("AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER")
		fmt.Println("LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,")
		fmt.Println("OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE")
		fmt.Println("SOFTWARE.")

		linha_0()
		time_3_sec()
		//confirma_enter()
		fmt.Printf("\rVoce entendeu ? ")
		CONFIRM = confirma_y_ou_n(CONFIRM)
		if CONFIRM {
			fmt.Printf("\n[voce] sim\n")
			time.Sleep(1 * time.Second)
			break
		} else {
			fmt.Printf("\n[voce] nao\n")
			time.Sleep(1 * time.Second)
			fmt.Printf("\n\ntudo bem, eu explico\n")
			time.Sleep(1 * time.Second)
		}

	}
	clear_screen()
}
