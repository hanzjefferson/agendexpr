package agendexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseField(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		bound       [2]uint8
		expectError error
		expectVal   uint64
	}{
		{
			name:        "anytime test",
			field:       "*",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   1<<60 - 1<<0,
		},
		{
			name:        "single test 1",
			field:       "1",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   1 << 1,
		},
		{
			name:        "single test 2",
			field:       "-1",
			bound:       [2]uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal:   0,
		},
		{
			name:        "single test 3",
			field:       "10000000000000000000000",
			bound:       [2]uint8{0, 59},
			expectError: ErrOverflow,
			expectVal:   0,
		},
		{
			name:        "single test 4",
			field:       "100",
			bound:       [2]uint8{0, 59},
			expectError: ErrNumberOutOfBound,
			expectVal:   0,
		},
		{
			name:        "part test 1",
			field:       "1,2,3",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   1<<1 | 1<<2 | 1<<3,
		},
		{
			name:        "part test 2",
			field:       "1,-2,3",
			bound:       [2]uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal:   0,
		},
		{
			name:        "step test 1",
			field:       "*/2",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   uint64(1<<60-1<<0) & stepMask[uint64](0, 59, 2),
		},
		{
			name:        "step test 2",
			field:       "2/2",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   uint64(1<<60-1<<2) & stepMask[uint64](2, 59, 2),
		},
		{
			name:        "step test 3",
			field:       "1-10/2",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   uint64(1<<11-1<<1) & stepMask[uint64](1, 10, 2),
		},
		{
			name:        "step test 4",
			field:       "1-10/-2",
			bound:       [2]uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal:   0,
		},
		{
			name:        "step test 5",
			field:       "-10/2",
			bound:       [2]uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal:   0,
		},
		{
			name:        "range test 1",
			field:       "1-5",
			bound:       [2]uint8{0, 59},
			expectError: nil,
			expectVal:   fillMask[uint64](1, 5),
		},
		{
			name:        "range test 2",
			field:       "1-5a",
			bound:       [2]uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal:   0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := parseField[uint64](test.field, test.bound)
			if test.expectError != nil {
				assert.ErrorIs(t, err, test.expectError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expectVal, res)
		})
	}
}

func BenchmarkParseField(b *testing.B) {
	bound := [2]uint8{0, 59}

	for b.Loop() {
		_, _ = parseField[uint64]("1-5,10/2,50-59", bound)
	}
}
