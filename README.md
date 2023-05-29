# bidichk - checks for dangerous unicode character sequences

[![Test Status](https://github.com/breml/bidichk/workflows/Go%20Matrix/badge.svg)](https://github.com/breml/bidichk/actions?query=workflow%3AGo%20Matrix) [![Go Report Card](https://goreportcard.com/badge/github.com/breml/bidichk)](https://goreportcard.com/report/github.com/breml/bidichk) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

bidichk finds dangerous unicode character sequences in Go source files.

## Considered dangerous unicode characters

The following unicode characters are considered dangerous:

* U+202A: LEFT-TO-RIGHT-EMBEDDING
* U+202B: RIGHT-TO-LEFT-EMBEDDING
* U+202C: POP-DIRECTIONAL-FORMATTING
* U+202D: LEFT-TO-RIGHT-OVERRIDE
* U+202E: RIGHT-TO-LEFT-OVERRIDE
* U+2066: LEFT-TO-RIGHT-ISOLATE
* U+2067: RIGHT-TO-LEFT-ISOLATE
* U+2068: FIRST-STRONG-ISOLATE
* U+2069: POP-DIRECTIONAL-ISOLATE

## Installation

Download `bidichk` from the [releases](https://github.com/breml/bidichk/releases) or get the latest version from source with:

```shell
go get github.com/breml/bidichk/cmd/bidichk
```

## Usage

### golangci-lint

[golangci-lint](https://golangci-lint.run) supports bidichk, so you can enable this linter and use it.

### Shell

Check everything:

```shell
bidichk ./...
```

### Enable only required unicode runes

If you run bidichk via golangci-lint look at [.golangci.example.yml](https://golangci-lint.run/usage/configuration/#config-file) for an example of the configuration.

Otherwise you can run bidichk with `--disallowed-runes` flag to specify the runes you consider harmful.

E.g. the following command considers only the `LEFT-TO-RIGHT-OVERRIDE` unicode rune as dangerous:

```shell
bidichk --disallowed-runes LEFT-TO-RIGHT-OVERRIDE ./...
```

For the full list of supported unicode runes [see above](#considered-dangerous-unicode-characters) or use

```shell
bidichk --help
```

## Inspiration

* ['Trojan Source' Bug Threatens the Security of All Code](https://krebsonsecurity.com/2021/11/trojan-source-bug-threatens-the-security-of-all-code/)
* [Trojan Source - Proofs-of-Concept](https://github.com/nickboucher/trojan-source)
* [Go proposal: disallow LTR/RTL characters in string literals?](https://github.com/golang/go/issues/20209)
* [cockroachdb/cockroach - PR: lint: add linter for unicode directional formatting characters](https://github.com/cockroachdb/cockroach/pull/72287)
* [Russ Cox - On “Trojan Source” Attacks](https://research.swtch.com/trojan)
