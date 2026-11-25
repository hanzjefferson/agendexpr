package agendexpr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpression(t *testing.T) {
	tests := []struct {
		name      string
		expr      string
		expectErr error
	}{
		{
			name: "success",
			expr: "* * * * * *",
		},
		{
			name:      "invalid format",
			expr:      "* * * * * * *",
			expectErr: ErrInvalidFormat,
		},
		{
			name:      "empty",
			expr:      "",
			expectErr: ErrEmpty,
		},
		{
			name:      "invalid digits",
			expr:      "-8",
			expectErr: ErrParseDigits,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Parse(test.expr)
			if test.expectErr != nil {
				assert.ErrorIs(t, err, test.expectErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExpression_Next(t *testing.T) {
	baseTime := time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		expr     string
		from     time.Time
		wantTime time.Time
	}{
		{
			name:     "any",
			expr:     "* * * * * *",
			from:     baseTime,
			wantTime: baseTime.Add(time.Second),
		},
		{
			name:     "next minute",
			expr:     "0 * * * * *",
			from:     baseTime.Add(time.Second * 30),
			wantTime: baseTime.Add(time.Second * 30).Truncate(time.Minute).Add(time.Minute),
		},
		{
			name:     "hourly step reaches last hour",
			expr:     "0 0 */1 * * *",
			from:     time.Date(2026, 7, 24, 22, 30, 0, 0, time.UTC),
			wantTime: time.Date(2026, 7, 24, 23, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.Next(test.from)
			assert.Equal(t, test.wantTime, got)
		})
	}
}

func TestExpression_DueMonth(t *testing.T) {
	tests := []struct {
		name string
		expr string
		of   time.Time
		want bool
	}{
		{
			name: "match",
			expr: "* * * * 7 *",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "mismatch",
			expr: "* * * * 8 *",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.DueMonth(test.of)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExpression_DueDay(t *testing.T) {
	tests := []struct {
		name string
		expr string
		of   time.Time
		want bool
	}{
		{
			name: "match",
			expr: "* * * 24 7 5",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "mismatch",
			expr: "* * * 25 7 *",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.DueDay(test.of)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExpression_DueHour(t *testing.T) {
	tests := []struct {
		name string
		expr string
		of   time.Time
		want bool
	}{
		{
			name: "match",
			expr: "* * 10 24 7 *",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "mismatch",
			expr: "* * 11 24 7 *",
			of:   time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.DueHour(test.of)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExpression_DueMinute(t *testing.T) {
	tests := []struct {
		name string
		expr string
		of   time.Time
		want bool
	}{
		{
			name: "match",
			expr: "* 30 10 24 7 *",
			of:   time.Date(2026, 7, 24, 10, 30, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "mismatch",
			expr: "* 0 10 24 7 *",
			of:   time.Date(2026, 7, 24, 10, 30, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.DueMinute(test.of)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExpression_Due(t *testing.T) {
	tests := []struct {
		name string
		expr string
		of   time.Time
		want bool
	}{
		{
			name: "match",
			expr: "30 15 14 * * *",
			of:   time.Date(2026, 1, 1, 14, 15, 30, 0, time.UTC),
			want: true,
		},
		{
			name: "mismatch",
			expr: "0 15 14 * * *",
			of:   time.Date(2026, 1, 1, 14, 15, 30, 0, time.UTC),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.Due(test.of)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestExpression_Countdown(t *testing.T) {
	baseTime := time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		expr string
		from time.Time
		want int64
	}{
		{
			name: "a second",
			expr: "* * * * * *",
			from: baseTime,
			want: 1,
		},
		{
			name: "a hour",
			expr: "0 0 * * * *",
			from: baseTime.Add(time.Minute * 30),
			want: 1800,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expr, err := Parse(test.expr)
			assert.NoError(t, err)

			got := expr.Countdown(test.from)
			assert.Equal(t, test.want, got)
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		_, _ = Parse("0 15 10 1 * *")
	}
}

func BenchmarkExpression_Next(b *testing.B) {
	expr, _ := Parse("0 0 1 1 * *")
	baseTime := time.Date(2026, 7, 24, 10, 0, 0, 0, time.UTC)

	for b.Loop() {
		_ = expr.Next(baseTime)
	}
}

func BenchmarkExpression_Due(b *testing.B) {
	expr, _ := Parse("30 15 14 * * *")
	checkTime := time.Date(2026, 7, 24, 14, 15, 30, 0, time.UTC)

	for b.Loop() {
		_ = expr.Due(checkTime)
	}
}
