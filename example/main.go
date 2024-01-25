package main

import (
	"github.com/Diegiwg/cli"
)

func helpCommand(ctx *cli.Context) error {
	println("Usage: " + ctx.App.Program + " <command> [arguments]")
	return nil
}

func main() {
	app := cli.NewApp()

	// app.SetDefaultRoutine(helpCommand)
	app.AddCommand(&cli.Command{
		Name: "help",
		Desc: "Show this Help Message",
		Exec: helpCommand,
	})

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
