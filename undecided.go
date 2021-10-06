package multiless

type (
	undecided struct{}
)

func (me undecided) LazySameLess(f func() (same bool, less bool)) Computation {
	panic("implement me")
}

func (me undecided) EagerSameLess(same, less bool) Computation {
	if same {
		return undecided{}
	}
	return decided(less)
}

// Sorts so that false comes before true.
func (me undecided) Bool(l, r bool) Computation {
	if l == r {
		return me
	} else {
		return decided(r)
	}
}

func (me undecided) Uint32(l, r uint32) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) Int64(l, r int64) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) Uint64(l, r uint64) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) Int(l, r int) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) CmpInt64(i int64) Computation {
	return me.EagerSameLess(i == 0, i < 0)
}

func (me undecided) Cmp(i int) Computation {
	return me.EagerSameLess(i == 0, i < 0)
}

func (me undecided) Uintptr(l, r uintptr) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) Less() bool {
	return false
}

func (me undecided) Ok() bool {
	return false
}

func (me undecided) LessOk() (less, ok bool) {
	return false, false
}

func (me undecided) MustLess() bool {
	panic("undecided")
}

func (me undecided) Float64(l, r float64) Computation {
	return me.EagerSameLess(l == r, l < r)
}

func (me undecided) Lazy(f func() Computation) Computation {
	return f()
}

func (me undecided) AndThen(then Computation) Computation {
	return then
}
