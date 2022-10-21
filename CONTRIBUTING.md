# Contributing

By participating to this project, you agree to abide our [code of conduct](/CODEOFCONDUCT.md).

## Setup your machine

`kiae` is written in [Go](https://golang.org/).

Prerequisites:

- `just`
- [Go 1.16+](https://golang.org/doc/install)

Clone `kiae` anywhere:

```sh
$ git clone git@github.com:kiaedev/kiae.git
```

Install the build and lint dependencies:

```sh
$ just install-dev-deps
```

A good way of making sure everything is all right is running the test suite:

```sh
$ just test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
$ just build
```

Which runs all the linters and tests.

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `kiae` fork and open a pull request against the
master branch.

## Financial contributions

We also welcome financial contributions in full transparency on our [open collective](https://opencollective.com/kiae).
Anyone can file an expense. If the expense makes sense for the development of the community, it will be "merged" in the ledger of our open collective by the core contributors and the person who filed the expense will be reimbursed.

## Credits

### Contributors

Thank you to all the people who have already contributed to kiae!
<a href="https://github.com/kiaedev/kiae/graphs/contributors"><img src="https://opencollective.com/kiae/contributors.svg?width=890" /></a>

### Backers

Thank you to all our backers! [[Become a backer](https://opencollective.com/kiae#backer)]

### Sponsors

Thank you to all our sponsors! (please ask your company to also support this open source project by [becoming a sponsor](https://opencollective.com/kiae#sponsor))
