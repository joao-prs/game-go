package main

import (
    "fmt"
	"time"
)

func panel_sobre_0() {
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