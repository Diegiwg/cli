# CLI Lib

Bem-vindo ao CLI Lib, uma biblioteca em Go para simplificar a criação de interfaces de linha de comando (CLIs). Esta biblioteca permite que você construa facilmente aplicativos CLI estruturados e adicione comandos de forma intuitiva.

## Instalação

Para começar a usar a CLI Lib, você pode instalá-la usando o seguinte comando:

```bash
go get -u github.com/Diegiwg/cli
```

## Exemplo de Uso

Aqui está um exemplo prático de como usar a CLI Lib para criar um simples calculadora CLI:

```go
package main

import (
 "errors"
 "strconv"
 "strings"

 "github.com/Diegiwg/cli"
)

func calc(ctx *cli.Context) error {
 if len(ctx.Args) < 2 {
  return errors.New("não foram fornecidos números suficientes")
 }

 a, err := strconv.Atoi(ctx.Args[0])
 if err != nil {
  return err
 }

 b, err := strconv.Atoi(ctx.Args[1])
 if err != nil {
  return err
 }

 op, ok := ctx.Flags["op"]
 if !ok || !strings.ContainsAny(op, "+-*/") {
  return errors.New("operação inválida")
 }

 switch op {
 case "+":
  {
   println(a + b)
  }
 case "-":
  {
   println(a - b)
  }
 case "*":
  {
   println(a * b)
  }
 case "/":
  {
   println(a / b)
  }
 }

 return nil
}

func main() {
 // Criar um novo aplicativo
 app := cli.NewApp()

 // Adicionar o comando de cálculo ao aplicativo
 app.AddCommand(&cli.Command{
  Name:  "calc",
  Desc:  "Calculadora Simples",
  Help:  "Esta é uma calculadora simples para somar, subtrair, multiplicar e dividir números.\n\tPasse os números como argumentos e a operação como uma flag.",
  Usage: "<a> <b> --op <operacao: + - / * >",
  Exec:  calc,
 })

 // Executar o aplicativo
 err := app.Run()
 if err != nil {
  println(err.Error())
 }
}
```

Neste exemplo, um aplicativo chamado "CalcApp" é criado com a descrição "Uma simples calculadora CLI". É adicionado um comando chamado "calc" que representa a calculadora, com uma descrição, ajuda, uso e uma função de execução associada. O comando pode ser chamado com argumentos numéricos e uma flag de operação.

## API da CLI Lib

### `cli.NewApp() *cli.App`

Cria um novo aplicativo CLI.

### `App.AddCommand(command Command)`

Adiciona um novo comando ao aplicativo. Um comando é uma estrutura que inclui campos como nome do comando, descrição, ajuda, modo de uso e uma função de execução.

### `cli.Command{Name string, Desc string, Help string, Usage string, Exec func(ctx *Context)}`

Estrutura que representa um comando na CLI. Os campos incluem o nome do comando, descrição, ajuda, modo de uso e uma função de execução.

### `cli.Context{App *App, Args []string, Flags map[string]string}`

Estrutura que é passada para a função de execução de um comando quando é chamado. Contém referência ao aplicativo, argumentos e flags identificados.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas, propor melhorias ou enviar pull requests.

## Licença

Este projeto é licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
