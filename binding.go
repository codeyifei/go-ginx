package ginx

type BindingType uint16

type BindingTypes []BindingType

func (b BindingTypes) Len() int {
	return len(b)
}

func (b BindingTypes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b BindingTypes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

const (
	BindingTypeUnknown BindingType = iota
	BindingTypeDefault
	BindingTypePath
	BindingTypeQuery
	BindingTypeJson
	BindingTypeHeader
)

func ParseBindingType(s string) BindingType {
	switch s {
	case "", "body":
		return BindingTypeDefault
	case "path", "uri":
		return BindingTypePath
	case "query", "form":
		return BindingTypeQuery
	case "json":
		return BindingTypeJson
	default:
		return BindingTypeUnknown
	}
}
