//go:build go1.27

package multiless

import "golang.org/x/exp/constraints"

// Compare applies the builtin ordering of l and r if the computation isn't decided yet. Unlike
// [Computation.Then], which takes the result of a three-way comparison, it takes the two values.
// Unlike [cmp.Compare], a NaN operand doesn't sort first: it decides the computation as not less.
func (me Computation) Compare[T constraints.Ordered](l, r T) Computation {
	return eagerOrdered(me, l, r)
}
