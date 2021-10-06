package multiless

type decided bool

func (d decided) EagerSameLess(same, less bool) Computation {
	return d
}

func (d decided) LessOk() (less, ok bool) {
	return bool(d), true
}

func (d decided) LazySameLess(f func() (same bool, less bool)) Computation {
	return d
}

func (d decided) CmpInt64(i int64) Computation {
	return d
}

func (d decided) Uint32(u uint32, u2 uint32) Computation {
	return d
}

func (d decided) Lazy(f func() Computation) Computation {
	return d
}

func (d decided) MustLess() bool {
	return bool(d)
}

func (d decided) Float64(f float64, f2 float64) Computation {
	return d
}

func (d decided) Uintptr(u uintptr, u2 uintptr) Computation {
	return d
}

func (d decided) Int64(i int64, i2 int64) Computation {
	return d
}

func (d decided) Cmp(i int) Computation {
	return d
}

func (d decided) Ok() bool {
	return true
}

func (d decided) Uint64(u uint64, u2 uint64) Computation {
	return d
}

func (d decided) Less() bool {
	return bool(d)
}

func (d decided) Int(i int, i2 int) Computation {
	return d
}

func (d decided) AndThen(computation Computation) Computation {
	return d
}

func (d decided) Bool(b bool, b2 bool) Computation {
	return d
}

var _ Computation = decided(false)
