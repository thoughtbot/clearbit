package main

import (
	"io"
	"net/url"
	"os"

	"github.com/codegangsta/cli"
	"github.com/thoughtbot/clearbit"
)

var prospectCommand = cli.Command{
	Name:      "prospect",
	Usage:     "fetch contacts for a company",
	ArgsUsage: "domain",
	Action:    prospect,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "Filter by first or last name (case-insensitive)"},
		cli.StringSliceFlag{Name: "title", Usage: "Filter by job title"},
	},
}

func prospect(ctx *cli.Context) {
	var (
		apiKey = apiKeyFromContext(ctx)
		domain = requiredArg(ctx, 0)
		name   = ctx.String("name")
		titles = ctx.StringSlice("title")
	)

	client := clearbit.NewClient(apiKey)

	res, err := client.Get(
		clearbit.ProspectorPersonSearchURL,
		url.Values{
			"domain":   []string{domain},
			"email":    []string{"true"},
			"name":     []string{name},
			"titles[]": titles,
		},
	)
	if err != nil {
		abort(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
