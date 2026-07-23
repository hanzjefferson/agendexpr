package agendexpr

type tokenMask interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Token[T tokenMask] struct {
	Value T
	Range [2]uint8
}

func ParseToken[T tokenMask](tok string, bounds [2]uint8) (Token[T], error) {
	val, r, err := parseToken[T](tok, bounds)
	if err != nil {
		return Token[T]{}, err
	}
	return Token[T]{
		Value: val,
		Range: r,
	}, nil
}