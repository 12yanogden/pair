package pairs

type Pair[F comparable, S comparable] struct {
	First  F
	Second S
}

func (p1 Pair[F, S]) Equals(p2 Pair[F, S]) bool {
	return p1.First == p2.First && p1.Second == p2.Second
}
