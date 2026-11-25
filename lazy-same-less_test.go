package multiless

import "testing"

func TestLazySameLessDecides(t *testing.T) {
	cases := []struct {
		same, less bool
		want       int
	}{{false, true, -1}, {false, false, 1}, {true, true, 0}, {true, false, 0}}
	for _, c := range cases {
		got := New().LazySameLess(func() (bool, bool) { return c.same, c.less })
		if got != New().EagerSameLess(c.same, c.less) {
			t.Errorf("LazySameLess(%v, %v) = %+v, disagrees with EagerSameLess", c.same, c.less, got)
		}
		if o := got.OrderingInt(); o != c.want {
			t.Errorf("LazySameLess(%v, %v).OrderingInt() = %v, want %v", c.same, c.less, o, c.want)
		}
	}
	// An already-decided computation doesn't call lazy.
	New().Bool(false, true).LazySameLess(func() (bool, bool) {
		t.Fatal("lazy called on decided computation")
		return false, false
	})
	// An undecided computation stays undecided when lazy reports same, so later comparisons apply.
	if got := New().LazySameLess(func() (bool, bool) { return true, true }).Bool(true, false).OrderingInt(); got != 1 {
		t.Errorf("later comparison after same LazySameLess = %v, want 1", got)
	}
}
