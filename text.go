package types

import (
	"fmt"
	"strings"
)

type byteMap map[byte]int

func (b byteMap) Len() int { return len(b) }

// func (b byteMap) Less(i, j int) bool { return b[i] < b[j] }

// LetterFrequency maps the frequency of letters in a text sample
func Frequency(s string) map[byte]int {
	m := make(map[byte]int, 255)

	b := []byte(s)

	var f int
	var ok bool

	for _, c := range b {
		if f, ok = m[c]; ok {
			m[c] = f + 1
		} else {
			m[c] = 1
		}
	}
	return m
}

func ShowFrequency(s string) string {
	sb := strings.Builder{}
	defer sb.Reset()

	m := Frequency(s)

	for k, v := range m {
		s := fmt.Sprintf("%2v: %d\n", k, v)
		sb.WriteString(s)
	}

	return sb.String()
}
