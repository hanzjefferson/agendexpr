package agendexpr

import (
	"errors"
	"strings"
)

var (
	aliases = map[string][]string{
		"yearly":    {"0", "0", "0", "1", "1", "*", "*"},
		"monthly":   {"0", "0", "0", "1", "*", "*", "*"},
		"weekly":    {"0", "0", "0", "*", "*", "0", "*"},
		"hourly":    {"0", "0", "*", "*", "*", "*", "*"},
		"minutely":  {"0", "*", "*", "*", "*", "*", "*"},
		"5minutes":  {"0", "*/5", "*", "*", "*", "*", "*"},
		"10minutes": {"0", "*/10", "*", "*", "*", "*", "*"},
		"30minutes": {"0", "*/30", "*", "*", "*", "*", "*"},
		"ticks":     {"*", "*", "*", "*", "*", "*", "*"},
	}

	dowLiteralReplacer = strings.NewReplacer(
		"sun", "0",
		"mon", "1",
		"tue", "2",
		"wed", "3",
		"thu", "4",
		"fri", "5",
		"sat", "6",
	)

	monthLiteralReplacer = strings.NewReplacer(
		"jan", "1",
		"feb", "2",
		"mar", "3",
		"apr", "4",
		"may", "5",
		"jun", "6",
		"jul", "7",
		"aug", "8",
		"sep", "9",
		"oct", "10",
		"nov", "11",
		"dec", "12",
	)

	ErrUnknownAlias      = errors.New("unknown alias for an expression")
	ErrInvalidFormat     = errors.New("invalid format")
	ErrEmptyExpr         = errors.New("empty expression not allowed")
)

func normalize(expr string) ([]string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, ErrEmptyExpr
	}

	expr = strings.ToLower(expr)
	// resolve '@' prefix
	if strings.HasPrefix(expr, "@") {
		alias := aliases[strings.TrimPrefix(expr, "@")]
		if alias == nil {
			return nil, ErrUnknownAlias
		}
		return alias, nil
	}

	toks := strings.Fields(expr)
	if len(toks) > 7 {
		return nil, ErrInvalidFormat
	}

	if len(toks) >= 5 {
		toks[4] = monthLiteralReplacer.Replace(toks[4])
	}
	if len(toks) >= 6 {
		toks[5] = dowLiteralReplacer.Replace(toks[5])
	}

	if len(toks) < 7 {
		complete := aliases["ticks"]

		copy(complete, toks)
		toks = complete
	}

	return toks, nil
}
