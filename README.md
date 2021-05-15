unescape
======================================================================

[![GoDoc](https://pkg.go.dev/badge/github.com/takumakei/go-unescape)](https://godoc.org/github.com/takumakei/go-unescape)

Unescape converts each 3-byte encoded substring of the form "%AB" into the
hex-decoded byte 0xAB, and 6-byte encoded substring of the form "%uABCD" into
the hex-decoded rune 0xABCD.
It does not convert the substring if any % is not the form above.
