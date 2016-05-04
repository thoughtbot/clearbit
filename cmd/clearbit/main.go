package main

import (
	"bytes"
	"encoding/json"
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
	app.Version = "0.1"

	app.Commands = []cli.Command{
		enrichCommand,
		prospectCommand,
	}

	app.Run(os.Args)
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

func requiredArgError(msg string) error {
	return cli.NewExitError(msg, 1)
}

func display(item interface{}) {
	data, _ := json.MarshalIndent(item, "", "  ")

	os.Stdout.Write(data)
	os.Stdout.Write([]byte("\n"))
}

func exitError(err error) error {
	return cli.NewExitError(err.Error(), 1)
}
