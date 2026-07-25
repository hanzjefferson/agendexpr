package agendexpr

// MaskType is a constraint that limits the underlying integer types usable
// by Mask. Only unsigned integer types that fit in 1–8 bytes are allowed.
type MaskType interface {
	~uint64 | ~uint32 | ~uint16 | ~uint8
}
