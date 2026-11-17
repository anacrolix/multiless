//go:build go1.27

package multiless

import "golang.org/x/exp/constraints"

// Applies builtin ordering if necessary.
func (me Computation) EagerOrdered[T constraints.Ordered](l, r T) Computation {
	return eagerOrdered(me, l, r)
}
