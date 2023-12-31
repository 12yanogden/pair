package pair

import (
	"testing"
)

func TestEquals(t *testing.T) {
	p1, p2 := Pair[string, string]{"a", "b"}, Pair[string, string]{"a", "b"}
	expected := true
	actual := p1.Equals(p2)

	if expected != actual {
		t.Fatalf("\nExpected\t%t, the pairs are equal\nActual:\t\t%t, the pairs are NOT equal", expected, actual)
	}
}

func TestNotEquals(t *testing.T) {
	p1, p2 := Pair[string, string]{"a", "b"}, Pair[string, string]{"a", "c"}
	expected := false
	actual := p1.Equals(p2)

	if expected != actual {
		t.Fatalf("\nExpected\t%t, the pairs are NOT equal\nActual:\t\t%t, the pairs are equal", expected, actual)
	}
}
