package semver

import (
	"strconv"
	"strings"
)

func Compare(a, b string) int {
	result := 0
	// split the string in arrays
	partsA := strings.Split(a, ".")
	partsB := strings.Split(b, ".")
	i := 0
	for (i < len(partsA) || i < len(partsB)) && result == 0 {
		intA := int64(0)
		intB := int64(0)
		if len(partsA) > i {
			intA, _ = strconv.ParseInt(partsA[i], 10, 64)
		}
		if len(partsB) > i {
			intB, _ = strconv.ParseInt(partsB[i], 10, 64)
		}
		if intA > intB {
			result = 1
		} else if intA < intB {
			result = -1
		}
		i++
	}
	// compare the two arrays
	return result
}
