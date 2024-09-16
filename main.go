package main

import (
	"ds/tree"
)

func main() {
	//str := "{s<s(s)s>sdf(we)s{d}w[e<f>w(e)f[s]d]ewf}"
	//items := stack.NewStack()
	//println(items.IsStable(str))

	t := tree.NewTree[int](1)
	t.AddLeft(2)
	t.AddRight(3)

	t.Left.AddLeft(4)
	t.Left.AddRight(5)
	t.Right.AddLeft(6)
	t.Right.AddRight(7)
	t.Stringer()
}
