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

func enrich(ctx *cli.Context) {
	var (
		apiKey = apiKeyFromContext(ctx)
		query  = requiredArg(ctx, 0)
		client = clearbit.NewClient(apiKey, nil)
	)

	item, err := enrichPersonOrCompany(client, query)
	if err != nil {
		abort(err)
	}
	display(item)
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
