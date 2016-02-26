package main

import (
	"io"
	"net/url"
	"os"
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
		token  = requiredArg(ctx, 0)
	)

	client := clearbit.NewClient(apiKey)

	endpoint, params := prepareLookupRequest(token)

	res, err := client.Get(endpoint, params)
	if err != nil {
		abort(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func prepareLookupRequest(token string) (string, url.Values) {
	if isEmail(token) {
		return clearbit.StreamingPersonSearchURL, url.Values{"email": {token}}
	}

	return clearbit.StreamingCompanySearchURL, url.Values{"domain": {token}}
}

func isEmail(s string) bool {
	return strings.Contains(s, "@")
}
