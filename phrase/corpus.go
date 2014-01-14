package phrase

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

type Corpus []string

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

func (c Corpus) Phrase() string {
	return c[rand.Intn(len(c))]
}
