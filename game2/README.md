
para testar o código rode
```sh
go run ./main.go ./funcoes_1.go ./panel_guia.go ./panel_menu.go
# ou
go run ./*.go
```

para converter em binário rode
```sh
go build -o my-game ./main.go ./funcoes_1.go ./panel_guia.go ./panel_menu.go  
# ou
go build -o my-game ./*.go
```

> contém muitos erros, e eu não tive tempo para corrigi-los