# bidichk - check for dangerous unicode character sequences

bidichk finds dangerous unicode character sequences in Go source files.

## Considered dangerous unicode characters

The following unicode characters are considered dangerous:

* U+202A: LEFT-TO-RIGHT-EMBEDDING:
* U+202B: RIGHT-TO-LEFT-EMBEDDING:
* U+202C: POP-DIRECTIONAL-FORMATTING:
* U+202D: LEFT-TO-RIGHT-OVERRIDE:
* U+202E: RIGHT-TO-LEFT-OVERRIDE:
* U+2066: LEFT-TO-RIGHT-ISOLATE:
* U+2067: RIGHT-TO-LEFT-ISOLATE:
* U+2068: FIRST-STRONG-ISOLATE:
* U+2069: POP-DIRECTIONAL-ISOLATE:

## Inspiration

* ['Trojan Source' Bug Threatens the Security of All Code](https://krebsonsecurity.com/2021/11/trojan-source-bug-threatens-the-security-of-all-code/)
* [Trojan Source - Proofs-of-Concept](https://github.com/nickboucher/trojan-source)
* [Go proposal: disallow LTR/RTL characters in string literals?](https://github.com/golang/go/issues/20209)
* [cockroachdb/cockroach - PR: lint: add linter for unicode directional formatting characters](https://github.com/cockroachdb/cockroach/pull/72287)
