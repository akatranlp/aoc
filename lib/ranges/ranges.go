package ranges

type Range struct {
	Down  int
	Up    int
	Valid bool
}

func NewRangeCount(from, count int) Range {
	return Range{from, from + count - 1, true}
}

func NewRange(down, up int) Range {
	return Range{down, up, true}
}

func (r *Range) InRange(v int) bool {
	return v >= r.Down && v <= r.Up
}

func (r *Range) CombineRanges(v *Range) *Range {
	if v.Down <= r.Down && v.Up >= r.Down ||
		r.Down <= v.Down && r.Up >= v.Down ||
		v.Up >= r.Up && v.Down <= r.Up ||
		r.Up >= v.Up && r.Down <= v.Up {
		r.Down = min(r.Down, v.Down)
		r.Up = max(r.Up, v.Up)
		v.Valid = false
	}
	return v
}

type states int

const (
	FullOuter states = iota
	FullInnerLeft
	FullInnerRight
	ConnectLeft
	ConnectRight
)

func (r *Range) RangeInRange(v *Range) states {
	if r.Up < v.Down || r.Down > v.Up {
		return FullOuter
	} else if r.Up > v.Down && r.Down < v.Up {
		return FullInnerLeft
	} else if v.Up > r.Down && v.Down < r.Up {
		return FullInnerLeft
	} else if r.Down < v.Down && r.Up > v.Down {
		return ConnectLeft
	} else if r.Up > v.Up && r.Down < v.Up {
		return ConnectRight
	}
	panic("unreachable")
}

func (r *Range) SplitRange(left int) Range {
	if r.Down > left || r.Up < left {
		return Range{}
	}
	v := NewRange(left+1, r.Up)
	r.Up = left
	return v
}
