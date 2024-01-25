package cli

import (
	"errors"
	"os"
	"strings"
	"unicode/utf8"
)

type Context struct {
	App   *App
	Flags map[string]string
	Args  []string
}

type Command struct {
	Name  string
	Desc  string
	Help  string
	Usage string
	Exec  func(ctx *Context) error
}

type App struct {
	Program         string
	Args            []string
	ArgsPtr         int
	DefaultCommand  func(ctx *Context) error
	SelectedCommand *Command
	Commands        map[string]*Command
}

func defaultHelpCommand(ctx *Context) error {
	if len(ctx.Args) != 0 {
		cmd, exists := ctx.App.Commands[ctx.Args[0]]
		if exists {
			println("Command:", "\n\t"+ctx.Args[0])
			println("Description:", "\n\t"+cmd.Help)
			println("Usage:", "\n\t"+ctx.App.Program, ctx.Args[0], cmd.Usage)
			return nil
		}

		println("Command not found:", ctx.Args[0], "\n")
	}

	println("Usage:", "\n\t"+ctx.App.Program, "<command> [arguments] [flags]")
	println("Commands:")
	for k, v := range ctx.App.Commands {
		println("\t"+k+":", v.Desc)
	}

	return nil
}

func defaultDumpCommand(ctx *Context) error {
	println("Dumping...")
	println("Program:", ctx.App.Program)
	println("Args (", len(ctx.Args), ")")
	for k, v := range ctx.Args {
		println("\t", k, "=", v)
	}
	println("Flags (", len(ctx.Flags), ")")
	for k, v := range ctx.Flags {
		println("\t", k, "=", v)
	}
	return nil
}

func (app *App) EnableDumpCommand() {
	app.AddCommand(&Command{
		Name:  "dump",
		Desc:  "Dump Tool",
		Help:  "Dump the current context",
		Usage: "",
		Exec:  defaultDumpCommand,
	})
}

func NewApp() *App {
	return &App{
		Program:         os.Args[0],
		Args:            os.Args[1:],
		DefaultCommand:  defaultHelpCommand,
		SelectedCommand: nil,
		Commands:        make(map[string]*Command),
	}
}

func (app *App) Run() error {
	// Add the Help command
	app.AddCommand(&Command{
		Name: "help",
		Desc: "Help Command",
		Help: "Show this help message",
		Exec: defaultHelpCommand,
	})

	// Parse the command
	err := app.ParseCommand()
	if err != nil && app.DefaultCommand == nil {
		return err
	}

	// Create the context
	ctx := &Context{
		App: app,
	}
	app.ParseArgsAndFlags(ctx)

	// Check if the SelectedCommand is Not nil
	if app.SelectedCommand != nil {
		return app.SelectedCommand.Exec(ctx)
	}

	// Check if the DefaultCommand is Not nil
	if app.DefaultCommand != nil {
		return app.DefaultCommand(ctx)
	}

	return errors.New("no command provided")
}

func (app *App) SetDefaultCommand(routine func(ctx *Context) error) {
	app.DefaultCommand = routine
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
	app.Args = app.Args[1:]
	app.SelectedCommand = cmd
	return nil
}

func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-") && utf8.RuneCountInString(arg) >= 2 ||
		strings.HasPrefix(arg, "--") && utf8.RuneCountInString(arg) >= 3
}

func removeFlagPrefix(arg string) string {
	if strings.HasPrefix(arg, "--") {
		return arg[2:]
	}

	return arg[1:]
}

func (app *App) ParseArgsAndFlags(ctx *Context) {
	args := []string{}
	flags := make(map[string]string)

	for i := 0; i < len(app.Args); i++ {
		arg := app.Args[i]
		if arg == "" {
			continue
		}

		if !isFlag(arg) {
			args = append(args, arg)
			continue
		}

		parts := strings.Split(arg, "=")
		parts[0] = removeFlagPrefix(parts[0])

		if len(parts) == 2 && parts[1] != "" {
			flags[parts[0]] = parts[1]
			continue
		}

		if i+1 < len(app.Args) && !isFlag(app.Args[i+1]) {
			flags[parts[0]] = app.Args[i+1]
			i++
			continue
		}

		flags[parts[0]] = ""
	}

	ctx.Args = args
	ctx.Flags = flags
}
