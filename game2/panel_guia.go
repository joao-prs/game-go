package main

import (
    "fmt"
	"os"
	"time"
	"bufio"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome string
	Idade int
	Vida int
}

type Jogador struct {
	Pessoa
}

func panel_create_pessoa(pessoa *Pessoa) Pessoa {
	// panel_create_pessoa() Pessoa
    clear_screen()
    fmt.Printf("\n\nQual seu nome?\n\n")

    // Lendo o nome da pessoa
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        pessoa.Nome = scanner.Text()
    }

    // Lendo a idade da pessoa
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("\n\nQual sua idade?\n\n")
    idade, _ := reader.ReadString('\n')
    idade = strings.TrimSpace(idade)
    pessoa.Idade, _ = strconv.Atoi(idade)

	return p
}
func panel_list_pessoa() Pessoa {
	var p Pessoa
	var j Jogador
	j.Nome = p.Nome
	j.Idade = p.Idade
	j.Vida = 100
	return j
}


func panel_introducao_0(p *Pessoa) {

	//var p Pessoa
	clear_screen()
	linha_0_efeito()
	//linha_0()
	fmt.Printf("Bem vindo, %s!\n\n", p.Nome)
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

func panel_manual_do_game() {
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