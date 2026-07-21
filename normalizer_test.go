package gronexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizer(t *testing.T) {
	tests := []struct {
		name        string
		expr        string
		expectError error
		expectVal   string
	}{
		{
			name: "shortcut test 1",
			expr: "*",
			expectError: nil,
			expectVal: "* * * * * * *",
		},
		{
			name: "shortcut test 2",
			expr: "* * *",
			expectError: nil,
			expectVal: "* * * * * * *",
		},
		{
			name: "shortcut test 3",
			expr: "* * * * * * *",
			expectError: nil,
			expectVal: "* * * * * * *",
		},
		{
			name: "invalid format test",
			expr: "* * * * * * * * *",
			expectError: ErrInvalidFormat,
			expectVal: "",
		},
		{
			name: "prefix test 1",
			expr: "@ticks",
			expectError: nil,
			expectVal: "* * * * * * *",
		},
		{
			name: "prefix test 2",
			expr: "@ti",
			expectError: ErrUnknownAlias,
			expectVal: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Normalize(test.expr)	
			if test.expectError != nil {
				assert.ErrorIs(t, err, test.expectError)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expectVal, res)
		})
	}
}
