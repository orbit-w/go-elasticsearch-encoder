package encoder

const (
	esRange = "range"
)

type RangeTerm struct {
	r    map[string]any
	name string
}

func (rt *RangeTerm) Name(name string) {
	rt.name = name
}

func (rt *RangeTerm) Range(en string, v any) {
	switch en {
	case "gte", "gt", "lte", "lt":
		if rt.r == nil {
			rt.r = map[string]any{}
		}
		rt.r[en] = v
	default:
	}
}

func (rt *RangeTerm) Length() int {
	return len(rt.r)
}

func (rt *RangeTerm) Json() map[string]any {
	return map[string]any{
		esRange: map[string]any{
			rt.name: rt.r,
		},
	}
}
