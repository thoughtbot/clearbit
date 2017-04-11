package main

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/thoughtbot/clearbit"
)

var enrichCommand = cli.Command{
	Name:      "enrich",
	Usage:     "enrich a person or company",
	ArgsUsage: "email or domain",
	Action:    enrich,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "combined", Usage: "Enrich and get both person and company data (email only)"},
	},
}

func enrich(ctx *cli.Context) error {
	var (
		apiKey = apiKeyFromContext(ctx)
		client = clearbit.NewClient(apiKey, nil)

		combined = ctx.Bool("combined")
		query    = ctx.Args().First()
	)

	if query == "" {
		return requiredArgError("Usage: clearbit enrich <email|domain>")
	}

	item, err := enrichPersonOrCompany(client, query, combined)
	if err != nil {
		return exitError(err)
	}

	display(item)
	return nil
}

func enrichPersonOrCompany(c *clearbit.Client, query string, combined bool) (interface{}, error) {
	if isEmail(query) {
		if combined {
			return c.Enrich(query)
		} else {
			return c.EnrichPerson(query)
		}
	}

	return c.EnrichCompany(query)
}

func isEmail(s string) bool {
	return strings.Contains(s, "@")
}
