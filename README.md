clearbit
========

[![GoDoc Reference](https://godoc.org/github.com/thoughtbot/clearbit?status.svg)][GoDoc]
[![Circle CI](https://circleci.com/gh/thoughtbot/clearbit.svg?style=shield)](https://circleci.com/gh/thoughtbot/clearbit)

`clearbit` is a client library and command-line interface
for the [Clearbit] API.

[Clearbit]: https://clearbit.com/

Install
-------

For OS X users, you can install and upgrade through the Homebrew package:

```
brew tap thoughtbot/formulae
brew update
brew install clearbit
```

For anyone with Go installed, you can install from the command-line:

```
go get -u github.com/thoughtbot/clearbit/cmd/clearbit
```

Use it in your Go project:

```go
import "github.com/thoughtbot/clearbit"
```

Usage from the command line
---------------------------

To use the `clearbit` command,
first store your [Clearbit API key][clearbit-api-key] in `~/.clearbit_key`.

Then use the subcommands to interact with the different
[Clearbit API endpoints][clearbit-api].

Get detailed information about a person from their email address:

```
$ clearbit enrich b@thoughtbot.com
```

Get data about a company from its domain:

```
$ clearbit enrich thoughtbot.com
```

Get contact details for a company:

```
$ clearbit prospect -title CEO -title COO thoughtbot.com
```

Since each command produces JSON as its output,
the `clearbit` command pairs nicely with [`jq`][jq] for processing.

Get company copy-pasteable company contacts:

```
$ clearbit prospect thoughtbot.com |
    jq '.[] | "\(.name.fullName) <\(.email)>"'
```

Run `clearbit help [subcommand]` for details on additional options.

  [clearbit-api]: https://clearbit.com/docs
  [clearbit-api-key]: https://dashboard.clearbit.com/keys
  [jq]: https://stedolan.github.io/jq/

Usage from Go
-------------

The `clearbit` package exposes types for interacting with the Clearbit API and
the data it returns.

```go
import "github.com/thoughtbot/clearbit"

client, _ := clearbit.NewClient(os.Getenv("CLEARBIT_API_KEY"))

bernerd, _ := client.EnrichPerson("b@thoughtbot.com")
thoughtbot, _ := client.EnrichCompany(bernerd.Employment.Domain)

prospects, _ := client.Prospect(clearbit.ProspectQuery{
  Domain: "thoughtbot.com",
  Titles: []string{"CEO", "CTO", "VP"},
})

for _, prospect := range prospects {
  prospectDetails, _ := client.EnrichPerson(prospect.Email)
}
```

For detailed API documentation, [read the go docs][GoDoc].

  [GoDoc]: https://godoc.org/github.com/thoughtbot/clearbit

Contributing
------------

We love pull requests from everyone.
By participating in this project,
you agree to abide by the thoughtbot [code of conduct].

[code of conduct]: https://thoughtbot.com/open-source-code-of-conduct

We expect everyone to follow the code of conduct
anywhere in thoughtbot's project codebases,
issue trackers, chatrooms, and mailing lists.

License
-------

`clearbit` is Copyright (c) 2016 thoughtbot, inc. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

About
-----

`clearbit` is maintained by Bernerd Schaefer.

![thoughtbot](https://thoughtbot.com/logo.png)

clearbit is maintained and funded by thoughtbot, inc.
The names and logos for thoughtbot are trademarks of thoughtbot, inc.

We love open source software!
See [our other projects][community]
or [hire us][hire] to help build your product.

[community]: https://thoughtbot.com/community?utm_source=github
[hire]: https://thoughtbot.com/hire-us?utm_source=github
