package agendexpr

import "math/bits"

type fieldMask interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Field[T fieldMask] struct {
	Value T
	Range [2]uint8
}

func ParseField[T fieldMask](tok string, bounds [2]uint8) (Field[T], error) {
	val, r, err := parseField[T](tok, bounds)
	if err != nil {
		return Field[T]{}, err
	}
	return Field[T]{
		Value: val,
		Range: r,
	}, nil
}

func (t Field[T]) Next(from uint8) int {
	if t.Value == 0 {
		return -1
	}

	var filtered T
	if from < 63 {
		mask := ^T(0) << (from + 1)
		filtered = t.Value & mask
	} else {
		filtered = 0
	}

	if filtered != 0 {
		return bits.TrailingZeros64(uint64(filtered))
	}

	return bits.TrailingZeros64(uint64(t.Value))
}

func (t Field[T]) Match(val uint8) bool {
	if val < t.Range[0] || val > t.Range[1] {
		return false
	}
	return (t.Value & (T(1) << val)) != 0
}

func (t Field[T]) List() []uint8 {
	var result []uint8
	v := uint64(t.Value)
	
	for v != 0 {
		tz := uint8(bits.TrailingZeros64(v))
		result = append(result, tz)
		v &= v - 1
	}
	
	return result
}