package multiless

func New() Computation {
	return undecided{}
}

// A helper for long chains of "less-than" comparisons, where later comparisons are only required if
// earlier ones haven't resolved the comparison.
type Computation interface {
	Bool(bool, bool) Computation
	Int64(int64, int64) Computation
	Cmp(int) Computation
	Ok() bool
	Uint64(uint64, uint64) Computation
	Less() bool
	Int(int, int) Computation
	AndThen(Computation) Computation
	Uintptr(uintptr, uintptr) Computation
	Float64(float64, float64) Computation
	MustLess() bool
	Lazy(func() Computation) Computation
	Uint32(uint32, uint32) Computation
	CmpInt64(int64) Computation
	LazySameLess(func() (same, less bool)) Computation
	LessOk() (less, ok bool)
	EagerSameLess(same, less bool) Computation
}
