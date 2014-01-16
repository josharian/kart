package phrase

import "testing"

// phrase is a dummy Phrase for testing
type phrase string

func (p phrase) Phrase() string { return string(p) }

func TestClean(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "a", want: "a"},
		{in: "  a \t ", want: "a"},
		{in: "☃s☃n☃o☃w☃m☃a☃n☃!☃", want: "snowman!"},
	}

	for _, tt := range cases {
		c := Clean(phrase(tt.in))
		have := c.Phrase()
		if have != tt.want {
			t.Errorf("clean %q want %q have %q", tt.in, tt.want, have)
		}
	}
}

func TestTruncate(t *testing.T) {
	cases := []struct {
		in     string
		length int
		want   string
	}{
		{in: "", length: 0, want: ""},
		{in: "", length: 1, want: ""},
		{in: "a", length: 0, want: ""},
		{in: "a", length: 1, want: "a"},
		{in: "a", length: 2, want: "a"},
		{in: "abcdef", length: 6, want: "abcdef"},
	}

	for _, tt := range cases {
		c := Truncate(phrase(tt.in), tt.length)
		have := c.Phrase()
		if have != tt.want {
			t.Errorf("truncate %q want %q have %q", tt.in, tt.want, have)
		}
	}
}
