package agendexpr

import (
	"strings"
)

var (
    aliases = map[string][6]string{
        "monthly":   [6]string(strings.Fields(MonthlyAlias)),
        "weekly":    [6]string(strings.Fields(WeeklyAlias)),
        "hourly":    [6]string(strings.Fields(HourlyAlias)),
        "minutely":  [6]string(strings.Fields(MinutelyAlias)),
        "5minutes":  [6]string(strings.Fields(FiveMinutesAlias)),
        "10minutes": [6]string(strings.Fields(TenMinutesAlias)),
        "30minutes": [6]string(strings.Fields(HalfMinutesAlias)),
        "ticks":     [6]string(strings.Fields(TicksAlias)),
    }
)

func normalizeMonth(s string) string {
	if len(s) == 3 {
		switch {
		case strings.EqualFold(s, "jan"):
			return "1"
		case strings.EqualFold(s, "feb"):
			return "2"
		case strings.EqualFold(s, "mar"):
			return "3"
		case strings.EqualFold(s, "apr"):
			return "4"
		case strings.EqualFold(s, "may"):
			return "5"
		case strings.EqualFold(s, "jun"):
			return "6"
		case strings.EqualFold(s, "jul"):
			return "7"
		case strings.EqualFold(s, "aug"):
			return "8"
		case strings.EqualFold(s, "sep"):
			return "9"
		case strings.EqualFold(s, "oct"):
			return "10"
		case strings.EqualFold(s, "nov"):
			return "11"
		case strings.EqualFold(s, "dec"):
			return "12"
		}
	}
	return s
}

func normalizeDow(s string) string {
	if len(s) == 3 {
		switch {
		case strings.EqualFold(s, "sun"):
			return "0"
		case strings.EqualFold(s, "mon"):
			return "1"
		case strings.EqualFold(s, "tue"):
			return "2"
		case strings.EqualFold(s, "wed"):
			return "3"
		case strings.EqualFold(s, "thu"):
			return "4"
		case strings.EqualFold(s, "fri"):
			return "5"
		case strings.EqualFold(s, "sat"):
			return "6"
		}
	}
	return s
}

func normalize(expr string) ([6]string, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return [6]string{}, ErrEmpty
	}

	// resolve '@' prefix
	if expr[0] == '@' {
		alias, ok := aliases[strings.ToLower(expr[1:])]
		if !ok {
			return [6]string{}, ErrUnknownAlias
		}
		return alias, nil
	}

	res := aliases["ticks"]
	var (
		start int = -1
		count int
	)
	for i := 0; i <= len(expr); i++ {
		space := i == len(expr) || expr[i] == ' ' || expr[i] == '\t'

		if space {
			if start != -1 {
				if count >= 6 {
					return [6]string{}, ErrInvalidFormat
				}

				tok := expr[start:i]

				switch count {
				case 4:
					tok = normalizeMonth(tok)
				case 5:
					tok = normalizeDow(tok)
				}

				res[count] = tok
				count++
				start = -1
			}
		} else {
			if start == -1 {
				start = i
			}
		}
	}

	return res, nil
}
