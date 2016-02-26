package main

import (
	"fmt"
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
		cli.StringSliceFlag{Name: "title", Usage: "Job title to filter by"},
		cli.BoolTFlag{Name: "email", Usage: "Include contact emails"},
	},
}

func prospect(ctx *cli.Context) {
	var (
		apiKey       = apiKeyFromContext(ctx)
		domain       = requiredArg(ctx, 0)
		titles       = ctx.StringSlice("title")
		includeEmail = ctx.BoolT("email")
	)

	client := clearbit.NewClient(apiKey)

	res, err := client.Get(
		clearbit.ProspectorPersonSearchURL,
		url.Values{
			"domain":   []string{domain},
			"email":    []string{fmt.Sprintf("%t", includeEmail)},
			"titles[]": titles,
		},
	)
	if err != nil {
		abort(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
