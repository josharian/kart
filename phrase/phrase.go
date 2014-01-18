// Package phrase provides phrase generation.
package phrase

import "strings"

// Source is a source of phrases.
type Source interface {
	Phrase() string
}

// Clean returns a Source that cleans phrases returned
// by s. Cleaning consists of trimming leading and
// trailing whitespace and removes non-ascii characters.
func Clean(s Source) Source {
	return &clean{src: s}
}

type clean struct {
	src Source
}

func (c *clean) Phrase() string {
	w := c.src.Phrase()
	w = strings.TrimSpace(w)
	asciiOnly := func(r rune) rune {
		if r > 127 {
			return -1
		}
		return r
	}
	w = strings.Map(asciiOnly, w)
	return w
}

// Truncate returns a Source that truncates phrases
// returned by s at a maximum length.
func Truncate(s Source, length int) Source {
	return &truncate{src: s, length: length}
}

type truncate struct {
	src    Source
	length int
}

func (t *truncate) Phrase() string {
	w := t.src.Phrase()
	if len(w) > t.length {
		w = w[:t.length]
	}
	return w
}
