package pairs

import "fmt"

type Pair[F comparable, S comparable] struct {
	First  F
	Second S
}

// Return true if the pair given is equal to the pair called. Else, false.
func (p1 Pair[F, S]) Equals(p2 Pair[F, S]) bool {
	return p1.First == p2.First && p1.Second == p2.Second
}

// Return values in pair
func (p Pair[F, S]) Deconstruct() (F, S) {
	return p.First, p.Second
}

// Return string representation of pair
func (p Pair[F, S]) String() string {
	return fmt.Sprintf("{%v, %v}", p.First, p.Second)
}
