package set

type Node[T comparable] struct {
	Value            T
	Par, Left, Right *Node[T]
	SubtreeSize      int
	Color            bool
}

func NewNode[T comparable](value T) *Node[T] {
	p := &Node[T]{
		Value:       value,
		SubtreeSize: 1,
	}
	return p
}
