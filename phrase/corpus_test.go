package phrase

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCorpus(t *testing.T) {
	cases := []struct {
		src  string
		want Corpus
		err  bool
	}{
		{src: "a", want: Corpus{"a"}},
		{src: "a\nb", want: Corpus{"a", "b"}},
		{src: "", err: true},
	}
	for _, tt := range cases {
		r := strings.NewReader(tt.src)
		c, err := NewCorpus(r)
		if tt.err {
			if err == nil {
				t.Errorf("wanted error parsing %q, got none", tt.src)
			}
			continue
		}
		if err != nil {
			t.Errorf("unexpected error parsing %q: %v", tt.src, err)
		}
		if !reflect.DeepEqual(c, tt.want) {
			t.Errorf("misparsed corpus %q, want %v have %v", tt.src, tt.want, c)
		}
	}
}
