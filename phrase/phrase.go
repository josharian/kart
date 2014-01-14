// Package phrase provides phrase generation.
package phrase

import "strings"

type Source interface {
	Phrase() string
}

func Clean(s Source) Source {
	return &clean{src: s}
}

func Shout(s Source) Source {
	return &shout{src: s}
}

type clean struct {
	src Source
}

func (c *clean) Phrase() string {
	w := c.src.Phrase()
	w = strings.TrimSpace(w)
	w = strings.Replace(w, "  ", " ", -1)
	asciiOnly := func(r rune) rune {
		if r > 127 {
			return -1
		}
		return r
	}
	w = strings.Map(asciiOnly, w)
	if len(w) > 15 {
		w = w[:15]
	}
	return w
}

type shout struct {
	src Source
}

func (s *shout) Phrase() string {
	return strings.ToUpper(s.src.Phrase())
}
