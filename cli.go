package cli

import (
	"errors"
	"os"
)

type Context struct {
	App   *App
	Args  []string
	Flags map[string]string
}

type Command struct {
	Name string
	Desc string
	Help string
	Exec func(ctx *Context) error
}

type App struct {
	Program         string
	Args            []string
	ArgsPtr         int
	DefaultRoutine  func(ctx *Context) error
	SelectedCommand *Command
	Commands        map[string]*Command
}

func NewApp() *App {
	return &App{
		Program:         os.Args[0],
		Args:            os.Args[1:],
		ArgsPtr:         0,
		DefaultRoutine:  nil,
		SelectedCommand: nil,
		Commands:        make(map[string]*Command),
	}
}

func (app *App) Run() error {
	// Parse the command
	err := app.ParseCommand()
	if err != nil && app.DefaultRoutine == nil {
		return err
	}

	// Create the context
	ctx := &Context{
		App: app,
	}

	// Check if the SelectedCommand is Not nil
	if app.SelectedCommand != nil {
		return app.SelectedCommand.Exec(ctx)
	}

	// Check if the DefaultRoutine is Not nil
	if app.DefaultRoutine != nil {
		return app.DefaultRoutine(ctx)
	}

	return errors.New("unknown error")
}

func (app *App) SetDefaultRoutine(routine func(ctx *Context) error) {
	app.DefaultRoutine = routine
}

func (app *App) AddCommand(cmd *Command) error {
	// Check if the command already exists
	if _, exists := app.Commands[cmd.Name]; exists {
		return errors.New("command already exists")
	}

	app.Commands[cmd.Name] = cmd
	return nil
}

func (app *App) RunCommand() error {
	return nil
}

func (app *App) ParseCommand() error {
	// Check if the Args is empty
	if len(app.Args) == 0 {
		return errors.New("no command provided")
	}

	// Check if the command exists
	cmd, exists := app.Commands[app.Args[0]]
	if !exists {
		return errors.New("command not found")
	}

	// Set the SelectedCommand
	app.ArgsPtr++
	app.SelectedCommand = cmd
	return nil
}
