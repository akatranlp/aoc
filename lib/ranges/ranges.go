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
