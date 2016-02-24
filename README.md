clearbit
========

`clearbit` is a command-line interface to the Clearbit API.

Install
-------

```
go get github.com/thoughtbot/clearbit/cmd/clearbit
```

Usage
-----

To use the `clearbit` command,
first store your [Clearbit API key][clearbit-api-key] in `~/.clearbit_key`.

Then use the subcommands to interact with the different
[Clearbit API endpoints][clearbit-api].

For example:

```
> clearbit prospect -title CEO -title COO thoughtbot.com
[
  {
    "id": "...",
    "name": {
      "givenName": "Chad",
      "familyName": "Pytel",
      "fullName": "Chad Pytel"
    },
    "title": "Founder and CEO",
    "email": "chad@thoughtbot.com",
    "verified": true
  },
  {
    "id": "...",
    "name": {
      "givenName": "Matt",
      "familyName": "Jankowski",
      "fullName": "Matt Jankowski"
    },
    "title": "COO",
    "email": "mjankowski@thoughtbot.com",
    "verified": true
  }
]
```

  [clearbit-api]: https://clearbit.com/docs
  [clearbit-api-key]: https://dashboard.clearbit.com/keys

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

clearbit is Copyright (c) 2016 thoughtbot, inc. It is free software,
and may be redistributed under the terms specified in the [LICENSE] file.

[LICENSE]: /LICENSE

About
-----

clearbit is maintained by Bernerd Schaefer.

![thoughtbot](https://thoughtbot.com/logo.png)

clearbit is maintained and funded by thoughtbot, inc.
The names and logos for thoughtbot are trademarks of thoughtbot, inc.

We love open source software!
See [our other projects][community]
or [hire us][hire] to help build your product.

[community]: https://thoughtbot.com/community?utm_source=github
[hire]: https://thoughtbot.com/hire-us?utm_source=github
