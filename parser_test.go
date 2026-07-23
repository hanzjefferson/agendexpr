package agendexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToken(t *testing.T) {
	tests := []struct {
		name   string
		token  string
		bounds []uint8
		expectError error
		expectVal uint64
	}{
		{
			name: "anytime test",
			token: "*",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: 1<<60-1<<0,
		},
		{
			name: "single test 1",
			token: "1",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: 1<<1,
		},
		{
			name: "single test 2",
			token: "-1",
			bounds: []uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal: 0,
		},
		{
			name: "single test 3",
			token: "10000000000000000000000",
			bounds: []uint8{0, 59},
			expectError: ErrOverflow,
			expectVal: 0,
		},
		{
			name: "single test 4",
			token: "100",
			bounds: []uint8{0, 59},
			expectError: ErrNumberOutOfBounds,
			expectVal: 0,
		},
		{
			name: "part test 1",
			token: "1,2,3",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: 1<<1 | 1<<2 | 1<<3,
		},
		{
			name: "part test 2",
			token: "1,-2,3",
			bounds: []uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal: 0,
		},
		{
			name: "step test 1",
			token: "*/2",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: uint64(1<<60 - 1<<0) & stepMask[uint64](0,59,2),
		},
		{
			name: "step test 2",
			token: "2/2",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: uint64(1<<60 - 1<<2) & stepMask[uint64](2,59,2),
		},
		{
			name: "step test 3",
			token: "1-10/2",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: uint64(1<<11 - 1<<1) & stepMask[uint64](1,10,2),
		},
		{
			name: "step test 4",
			token: "1-10/-2",
			bounds: []uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal: 0,
		},
		{
			name: "step test 5",
			token: "-10/2",
			bounds: []uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal: 0,
		},
		{
			name: "range test 1",
			token: "1-5",
			bounds: []uint8{0, 59},
			expectError: nil,
			expectVal: fillMask[uint64](1,5),
		},
		{
			name: "range test 2",
			token: "1-5a",
			bounds: []uint8{0, 59},
			expectError: ErrParseDigits,
			expectVal: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := parseToken[uint64](test.token, test.bounds)
			if test.expectError != nil {
				assert.ErrorIs(t, err, test.expectError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expectVal, res)
		})
	}
}
