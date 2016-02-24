package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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
		{
			Name:      "prospect",
			Usage:     "fetch contacts for a company",
			ArgsUsage: "domain",
			Action:    prospectSearch,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: "title", Usage: "Job title to filter by"},
				cli.BoolTFlag{Name: "email", Usage: "Include contact emails"},
			},
		},
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

func abort(reason interface{}) {
	fmt.Printf("ERROR: %s\n", reason)
	os.Exit(1)
}

func prospectSearch(ctx *cli.Context) {
	domain := ctx.Args().First()
	if domain == "" {
		abort("domain required")
	}

	req, err := http.NewRequest("GET", "https://prospector.clearbit.com/v1/people/search", nil)
	if err != nil {
		abort(err)
	}

	req.SetBasicAuth(ctx.GlobalString("api-key"), "")

	req.URL.RawQuery = url.Values{
		"domain":   []string{domain},
		"email":    []string{fmt.Sprintf("%b", ctx.BoolT("email"))},
		"titles[]": ctx.StringSlice("title"),
	}.Encode()

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		abort(err)
	}

	io.Copy(os.Stdout, res.Body)
}
