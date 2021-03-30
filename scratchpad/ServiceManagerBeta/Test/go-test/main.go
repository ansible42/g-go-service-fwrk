package main

import (
	"fmt"
	"github.com/alecthomas/kong"
)

type Globals struct {
	Version VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

// Support verssion info
func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

//All the commands supported by this app
type CLI struct {
	Globals          // Always include the global options
	Count   CountCmd `cmd count: "Display a countdowns" `
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("scratchpad"),
		kong.Description("Me learning how to use kong"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
