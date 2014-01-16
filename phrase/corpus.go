package phrase

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

// Corpus is a list of phrases to select from.
type Corpus []string

// NewCorpus reads a corpus from r. Phrases
// in r are separated by newlines.
func NewCorpus(r io.Reader) (Corpus, error) {
	c := make(Corpus, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c = append(c, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(c) == 0 {
		return nil, fmt.Errorf("at least one line required")
	}
	return c, nil
}

// Phrase returns a randomly selected phrase from c.
func (c Corpus) Phrase() string {
	return c[rand.Intn(len(c))]
}
