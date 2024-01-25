package main

import (
	"strings"

	"github.com/Diegiwg/cli"
)

func main() {
	app := cli.NewApp()

	app.SetDefaultCommand(func(ctx *cli.Context) error {
		println(strings.Join(ctx.Args, " "))
		return nil
	})

	app.EnableDumpCommand()

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
