# Monitoramento Glô :mag_right:

## Feature do projeto que aprendi na Alura
Verificação se os sites selecionados da Glô estão online ou fora do ar. Todas as informações coletadas serão armazenadas no log e apresentarão datas e horários correspondentes de monitoramento e os respectivos status.

## Tecnologias utilizadas
* VS Code
* Golang (Go)
* Package fmt
* Package net/http
* Package os
* Package time
* Package bufio
* Package io
* Package strconv
* Package strings

## Ferramentas
Foi instalado a linguagem <a href="https://go.dev/doc/install" target="_blank"> Go</a> no Windows com a versão 1.22.2

Foi instalado a extensão Go no <a href="https://code.visualstudio.com/download" target="_blank"> VS Code</a> com a versão 0.44.0

Layouts predefinidos para o uso em <a href="https://go.dev/src/time/format.go" target="_blank"> Time.Format</a> com a conversão de data e hora específicos.

## Rodando o projeto
- Comando para rodar no terminal:
```
go run main.go
```

- Parar de rodar no terminal: clicar nas teclas de `"Ctrl"` e `"C"`.

- Clicando na opção 1 do menu: se iniciará o monitoramento de todos os sites salvos no arquivo `websites.txt`. E ao mesmo tempo, o arquivo `logs.txt` irá guardar todos os registros, incluindo informações de data completa, horário, url e o status.

- Clicando na opção 2 do menu: irá imprimir os logs que inicialmente retornam um array de bytes, mas em seguida foi convertido em string para exibir os logs no terminal.

- Clicando na opção 3 do menu: sairá do programa. 


## Status do projeto
:heavy_check_mark: Aplicação finalizada.
