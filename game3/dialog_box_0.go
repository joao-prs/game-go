package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	\033[32m verde
	\033[96m ciano
	\033[0m retorna
*/

func dialog_box_0_scene_0() {
	fmt.Println("\033[96mVoce entra em uma loja interessante, mas nao tem dinheiro")
	fmt.Println("para nada, então você ve um mendigo jogando dados sozinho, ele")
	fmt.Println("parece se divertir com aquilo. Ele tambem tem uma mochila bem")
	fmt.Println("desbotada, mas bem fechada e sem sinais de rasgados e avarias.")
	fmt.Println("ele o percebe o encarando e então o chama para jogar dados\033[0m\n")
}

func dialog_box_0_scene_1() {
	fmt.Println("\033[32mMendigo\033[0m")
	fmt.Println("- Voce parece interessado em mim ou no que eu estou fazendo.")
	fmt.Println("que tal jogar com estes dados, se voce acertar o valor de um")
	fmt.Println("deles depois que eu chaqualhar o copo com eles dentro, lhe dou")
	fmt.Println("10 reais, se voce errar, eu quero 20 reais.")
	time.Sleep(1 * time.Second)
	fmt.Println("\033[96mVoce não tem dinheiro, mas não vê problema em jogar com ele,")
	fmt.Println("perceba que ele tem dois dados D8, então a probabilidade de")
	fmt.Println("você acertar é baixa, mas não nula. Alias aquela mochila é atraente,")
	fmt.Println("talvez haja algo dentro...\033[0m\n")
}

func dialog_box_0_scene_2() {
	fmt.Println("\033[32mMendigo\033[0m")
	fmt.Println("- Vamos lá!")
	time.Sleep(1 * time.Second)
	fmt.Println("\033[96mOs dados começam a rolar dentro do copo\033[0m\n\n")
	time.Sleep(1 * time.Second)
}

func dialog_box_0_scene_3() {
	fmt.Println("\033[32mMendigo\033[0m")
	time.Sleep(1 * time.Second)
	fmt.Println("- Vamos, escolha um valor entre 1 e 8.\n")

	var userChoice int

	for {
		// Solicita ao usuário para escolher um número entre 1 e 20
		//fmt.Print("Escolha um número entre 1 e 20: ")
		_, err := fmt.Scan(&userChoice)

		// Verifica se houve erro na entrada do usuário (como uma letra)
		if err != nil {
			fmt.Println("Entrada inválida! Por favor, escolha um número entre 1 e 8.")
			// Limpa o buffer de entrada
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		// Verifica se a escolha do usuário está dentro do intervalo válido
		if userChoice < 1 || userChoice > 8 {
			fmt.Println("Número fora do intervalo! Por favor, escolha um número entre 1 e 8.")
			continue
		}

		// Se a entrada for válida, sai do loop
		break
	}

	rand.Seed(time.Now().UnixNano()) // Inicializa o gerador de números aleatórios com um seed baseado no tempo atual

	randomValue := rand.Intn(8) + 1 // Gera um número aleatório entre 1 e 20

	fmt.Print("\n\033[96mpausa dramatica...\033[0m\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("\nO dado parou em %d\n", randomValue)

	if randomValue == userChoice {
		time.Sleep(1 * time.Second)
		fmt.Println("você deu sorte.")
		time.Sleep(2 * time.Second)
		fmt.Println("\033[32mMendigo\033[0m")
		fmt.Println("- Bom, eu não tenho dinheiro não, então não vou lhe pagar")
		time.Sleep(3 * time.Second)
		fmt.Println("\033[32mVocê\033[0m")
		fmt.Println("- Bom, eu não tambem não tenho, então estamos quites")
		time.Sleep(1 * time.Second) // saindo do  dialog_box_0_scene_3()
	} else {
		time.Sleep(1 * time.Second)
		fmt.Println("Você se deu mal! hahahahahahahahahahah")
		time.Sleep(2 * time.Second)
		fmt.Println("\033[32mMendigo\033[0m")
		fmt.Println("- Agora me paga o dinheiro! Me dê meus 20 reais!!")
		time.Sleep(3 * time.Second)
		fmt.Println("\033[32mVocê\033[0m")
		fmt.Println("- Sai fora maluco, não tenho dinheiro não.")
		time.Sleep(3 * time.Second)
		fmt.Println("\033[32mMendigo\033[0m")
		fmt.Println("- tic... Eu tambem não tenho, então deixa pra la,")
		fmt.Println("mesmo que tivesse, eu não daria pra você.")
		time.Sleep(1 * time.Second) // saindo do  dialog_box_0_scene_3()
	}
}

func dialog_box_0_scene_4() {
	fmt.Println("\033[32mMendigo\033[0m")
	fmt.Println("- Foi um bom jogo, deu para me distrair, agora dá o fora.")
	time.Sleep(1 * time.Second)
}
