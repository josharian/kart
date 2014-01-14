package phrase

import "math/rand"

const (
	DefaultRandChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 -./~"
)

type Rand struct {
	Chars  string
	Length int
}

func (r *Rand) Phrase() string {
	if r.Chars == "" {
		r.Chars = DefaultRandChars
	}
	nChars := len(r.Chars)
	var b []byte
	for i := 0; i < r.Length; i++ {
		b = append(b, r.Chars[rand.Intn(nChars)])
	}
	return string(b)
}
