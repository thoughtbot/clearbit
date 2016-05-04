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
}

func enrich(ctx *cli.Context) error {
	var (
		apiKey = apiKeyFromContext(ctx)
		client = clearbit.NewClient(apiKey, nil)

		query = ctx.Args().First()
	)

	if query == "" {
		return requiredArgError("Usage: clearbit enrich <email|domain>")
	}

	item, err := enrichPersonOrCompany(client, query)
	if err != nil {
		return exitError(err)
	}

	display(item)
	return nil
}

func enrichPersonOrCompany(c *clearbit.Client, query string) (interface{}, error) {
	if isEmail(query) {
		return c.EnrichPerson(query)
	}

	return c.EnrichCompany(query)
}

func isEmail(s string) bool {
	return strings.Contains(s, "@")
}
