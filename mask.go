package agendexpr

import "math/bits"

type Mask[T MaskType] struct {
	raw T
}

func (m Mask[T]) Raw() T {
	return m.raw
}

func (m Mask[T]) Next(from uint8) uint8 {
	var filtered T
	if from < 63 {
		mask := ^T(0) << (from + 1)
		filtered = m.raw & mask
	} else {
		filtered = 0
	}

	if filtered != 0 {
		return uint8(bits.TrailingZeros64(uint64(filtered)))
	}

	return uint8(bits.TrailingZeros64(uint64(m.raw)))
}

func (m Mask[T]) Match(val uint8) bool {
	return (m.raw & (T(1) << val)) != 0
}

func (m Mask[T]) List() []uint8 {
	var result []uint8
	v := uint64(m.raw)
	
	for v != 0 {
		tz := uint8(bits.TrailingZeros64(v))
		result = append(result, tz)
		v &= v - 1
	}
	
	return result
}