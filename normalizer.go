package gronexpr

import (
	"errors"
	"strings"
)

var (
	aliases = map[string]string{
		"yearly":    "0 0 0 1 1 * *",
		"monthly":   "0 0 0 1 * * *",
		"weekly":    "0 0 0 * * 0 *",
		"hourly":    "0 0 * * * * *",
		"minutely":  "0 * * * * * *",
		"5minutes":  "0 */5 * * * * *",
		"10minutes": "0 */10 * * * * *",
		"30minutes": "0 */30 * * * * *",
		"ticks":      "* * * * * * *",
	}

	literalReplacer = strings.NewReplacer(
	"sun", "0",
	"mon", "1",
	"tue", "2",
	"wed", "3",
	"thu", "4",
	"fri", "5",
	"sat", "6",
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

	ErrUnknownAlias = errors.New("unknown alias for an expression")
	ErrInvalidFormat = errors.New("invalid format")
	ErrEmptyExpr = errors.New("empty expression not allowed")
)

func Normalize(expr string) (string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return "", ErrEmptyExpr
	}

	expr = strings.ToLower(expr)
	// resolve '@' prefix
	if strings.HasPrefix(expr, "@") {
		alias := aliases[strings.TrimPrefix(expr, "@")]
		if alias == "" {
			return "", ErrUnknownAlias
		}
		return alias, nil
	}

	expr = literalReplacer.Replace(expr)

	tok := strings.Fields(expr)
	if len(tok) > 7 {
		return "", ErrInvalidFormat
	}
	if len(tok) < 7 {
		complete := strings.Fields(aliases["ticks"])
	
		copy(complete, tok)
		tok = complete
	}

	expr = strings.Join(tok, " ")

	return expr, nil
}
