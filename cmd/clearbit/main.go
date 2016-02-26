package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/codegangsta/cli"
)

func main() {
	var clearbitKey string

	app := cli.NewApp()
	app.Name = "clearbit"
	app.Usage = "interact with the Clearbit API"
	app.Before = clearbitKeyLoader(&clearbitKey)
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "api-key", Usage: "clearbit API key (default ~/.clearbit_key)", Destination: &clearbitKey},
	}

	app.Commands = []cli.Command{
		enrichCommand,
		prospectCommand,
	}
	app.RunAndExitOnError()
}

func clearbitKeyLoader(key *string) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		if ctx.String("api-key") != "" {
			return nil
		}

		currentUser, err := user.Current()
		if err != nil {
			return err
		}

		keyFile := path.Join(currentUser.HomeDir, ".clearbit_key")

		data, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}

		data = bytes.TrimSpace(data)
		*key = string(data)
		return nil
	}
}

func apiKeyFromContext(ctx *cli.Context) string {
	return ctx.GlobalString("api-key")
}

func requiredArg(ctx *cli.Context, n int) string {
	arg := ctx.Args().Get(n)

	if arg == "" {
		cli.ShowSubcommandHelp(ctx)
		os.Exit(1)
	}

	return arg
}

func abort(reason interface{}) {
	fmt.Printf("ERROR: %s\n", reason)
	os.Exit(1)
}
