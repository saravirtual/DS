package tree

import "fmt"

type tree[T any] struct {
	top   *tree[T]
	Left  *tree[T]
	Right *tree[T]
	value T
	level uint
}

func NewTree[T any](v T) *tree[T] {
	return &tree[T]{
		Left: nil, Right: nil, value: v, level: 0,
	}
}

func (t *tree[T]) AddLeft(v T) {
	t.Left = NewTree(v)
	t.Left.top = t
	t.Left.level = t.level + 1
}

func (t *tree[T]) AddRight(v T) {
	t.Right = NewTree(v)
	t.Right.top = t
	t.Right.level = t.level + 1
}

func (t *tree[T]) Stringer() {
	fmt.Print(t.value, " ")
	if t.Left != nil {
		t.Left.level++
		if t.top == nil {
			fmt.Print("\n", indent(t.level+1))
		} else {
			fmt.Print("\n", indent(t.level))
		}
		t.Left.Stringer()
	}
	if t.Right != nil {
		t.Right.level++
		if t.top == nil {
			fmt.Print("\n", indent(t.level+1))
		} else {
			fmt.Print("\n", indent(t.level))
		}
		t.Right.Stringer()
	}
}

func indent(times uint) string {
	ind := ""
	for i := uint(0); i < times; i++ {
		//ind += fmt.Sprint("\t")
		ind += fmt.Sprint("  ")
	}
	return ind
}
