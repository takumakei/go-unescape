package unescape

import (
	"bytes"
	"regexp"
)

// Unescape converts each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB, and 6-byte encoded substring of the form "%uABCD"
// into the hex-decoded rune 0xABCD.
// It does not convert the substring if any % is not the form above.
func Unescape(s string) string {
	b := bytes.NewBuffer(make([]byte, 0, len(s)))
	i := 0
	for _, v := range re.FindAllStringSubmatchIndex(s, -1) {
		j := v[0]
		b.WriteString(s[i:j])
		if k := v[3]; k >= 0 {
			b.WriteString(unescape_uXXXX(s, j, k))
			i = k
		} else {
			k := v[5]
			b.WriteString(unescape_XX(s, j, k))
			i = k
		}
	}
	b.WriteString(s[i:])
	return b.String()
}

var re = regexp.MustCompile(`((?:%u[0-9a-fA-F]{4})+)|((?:%[0-9a-fA-F]{2})+)`)

func unescape_uXXXX(s string, i, j int) string {
	a := make([]rune, 0, (j-i)/6)
	for i += 2; i < j; i += 6 {
		v := decode_XXXX(s[i : i+4])
		a = append(a, rune(v))
	}
	return string(a)
}

func unescape_XX(s string, i, j int) string {
	a := make([]byte, 0, (j-i)/3)
	for i += 1; i < j; i += 3 {
		v := decode_XX(s[i : i+2])
		a = append(a, byte(v))
	}
	return string(a)
}

func decode_XXXX(s string) rune {
	var n rune
	for _, c := range []byte(s) {
		n = n*16 + rune(decode(c))
	}
	return n
}

func decode_XX(s string) byte {
	var n byte
	for _, c := range []byte(s) {
		n = n*16 + decode(c)
	}
	return n
}

func decode(c byte) byte {
	var d byte
	switch {
	case 'a' <= c && c <= 'f':
		d = c - 'a' + 10
	case 'A' <= c && c <= 'F':
		d = c - 'A' + 10
	default:
		d = c - '0'
	}
	return d
}
