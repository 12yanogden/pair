package pairs

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

func TestDeconstruct(t *testing.T) {
	p := Pair[string, string]{"a", "b"}
	expected1, expected2 := "a", "b"
	actual1, actual2 := p.Deconstruct()

	if expected1 != actual1 || expected2 != actual2 {
		t.Fatalf("\nExpected:\t%s, %s\nActual:\t\t%s, ,%s", expected1, expected2, actual1, actual2)
	}
}

func TestString(t *testing.T) {
	p := Pair[string, string]{"a", "b"}
	expected := "{a, b}"
	actual := p.String()

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s", expected, actual)
	}
}
