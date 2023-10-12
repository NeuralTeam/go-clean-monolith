package rotator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Byte size size suffixes.
const (
	B  uint64 = 1
	KB uint64 = 1 << (10 * iota)
	MB
	GB
)

// Used to convert user input to ByteSize
var stringUnitMap = map[string]uint64{
	"B":  B,
	"KB": KB,
	"MB": MB,
	"GB": GB,
}

func ParseStringSize(s string) (uint64, error) {
	split := make([]string, 0)
	for i, r := range s {
		if !unicode.IsDigit(r) {
			split = append(split, strings.TrimSpace(s[:i]))
			split = append(split, strings.TrimSpace(s[i:]))
			break
		}
	}
	if len(split) != 2 {
		return 0, errors.New("len error: != 2")
	}
	unit, ok := stringUnitMap[split[1]]
	if !ok {
		return 0, errors.New("len error: != 2")
	}
	value, err := strconv.ParseUint(split[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return value * unit, nil
}
