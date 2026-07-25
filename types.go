package agendexpr

type MaskType interface {
    ~uint64 | ~uint32 | ~uint16 | ~uint8
}