package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Diegiwg/cli"
)

func calc(ctx *cli.Context) error {
	if len(ctx.Args) < 2 {
		return errors.New("not enough of numbers provided")
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
		return errors.New("invalid operation")
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
	app := cli.NewApp()
	app.EnableDumpCommand()

	app.AddCommand(&cli.Command{
		Name:  "calc",
		Desc:  "Simple Calculator",
		Help:  "This is a simple calculator to add, subtract, multiply and divide numbers.\n\tPass the numbers as arguments and the operation as a flag.",
		Usage: "<a> <b> --op <operation: + - / * >",
		Exec:  calc,
	})

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
