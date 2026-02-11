package types

import (
	"strconv"
	"strings"
)

func parseClockTypeStr(clockTypeStr string) string {
	clockWoPT := strings.Replace(clockTypeStr, "PT", "", 1)
	clockParsedM := strings.Replace(clockWoPT, "M", ":", 1)
	return strings.Replace(clockParsedM, "S", "", 1)
}

func parseNumFromParsedClockTypeStr(parsedClockTypeStr string) (float64, bool) {
	num, err := strconv.ParseFloat(parsedClockTypeStr, 64)
	if err != nil {
		return 0, false
	}

	return num, true
}
