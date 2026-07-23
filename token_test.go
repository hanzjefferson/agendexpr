package agendexpr

import (
	"reflect"
	"testing"
)

func TestToken_Match(t *testing.T) {
	tests := []struct {
		name  string
		token Token[uint64]
		val   uint8
		expect  bool
	}{
		{
			name:  "matched",
			token: Token[uint64]{Value: (1 << 5) | (1 << 10), Range: [2]uint8{0, 59}},
			val:   5,
			expect:  true,
		},
		{
			name:  "not matched",
			token: Token[uint64]{Value: (1 << 5) | (1 << 10), Range: [2]uint8{0, 59}},
			val:   6,
			expect:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.token.Match(tt.val); got != tt.expect {
				t.Errorf("Token.Match(%d) = %v, expect %v", tt.val, got, tt.expect)
			}
		})
	}
}

func TestToken_List(t *testing.T) {
	tests := []struct {
		name  string
		token Token[uint64]
		expect  []uint8
	}{
		{
			name:  "empty",
			token: Token[uint64]{Value: 0},
			expect:  nil,
		},
		{
			name:  "only one",
			token: Token[uint64]{Value: 1 << 5},
			expect:  []uint8{5},
		},
		{
			name:  "many",
			token: Token[uint64]{Value: (1 << 1) | (1 << 15) | (1 << 50)},
			expect:  []uint8{1, 15, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.token.List()
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("Token.List() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestToken_Next(t *testing.T) {
	tests := []struct {
		name  string
		token Token[uint64]
		from  uint8
		expect  int
	}{
		{
			name:  "zero value",
			token: Token[uint64]{Value: 0},
			from:  5,
			expect:  -1,
		},
		{
			name:  "next value",
			token: Token[uint64]{Value: (1 << 5) | (1 << 10)},
			from:  5,
			expect:  10,
		},
		{
			name:  "next-of value",
			token: Token[uint64]{Value: (1 << 2) | (1 << 50)},
			from:  10,
			expect:  50,
		},
		{
			name:  "wrap",
			token: Token[uint64]{Value: (1 << 5) | (1 << 10)},
			from:  15,
			expect:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.token.Next(tt.from); got != tt.expect {
				t.Errorf("Token.Next(%d) = %v, expect %v", tt.from, got, tt.expect)
			}
		})
	}
}