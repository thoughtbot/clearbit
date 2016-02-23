Contributing
============

We love pull requests from everyone.
By participating in this project,
you agree to abide by the thoughtbot [code of conduct].

  [code of conduct]: https://thoughtbot.com/open-source-code-of-conduct

We expect everyone to follow the code of conduct
anywhere in thoughtbot's project codebases,
issue trackers, chatrooms, and mailing lists.

Fork the repo.

Get a working [Go installation],
and clone the project into your [Go work environment]
(that is, `$GOPATH/src/github.com/thoughtbot/$(REPO_NAME)`).

  [Go installation]: http://golang.org/doc/install
  [Go work environment]: http://golang.org/doc/code.html

Run `bin/setup` to install the project's dependencies.

If you add or update a dependency,
run `godep save ./...` to vendor the changes.

To test the `$(REPO_NAME)` package, run `go test ./...`.

Make your change, with new passing tests.

Push to your fork. Write a [good commit message][commit]. Submit a pull request.

  [commit]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html

Others will give constructive feedback.
This is a time for discussion and improvements,
and making the necessary changes will be required before we can
merge the contribution.
