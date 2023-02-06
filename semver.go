package semver

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Compare(a, b string) (int, error) {
	if len(strings.TrimSpace(a)) == 0 || len(strings.TrimSpace(b)) == 0 {
		return 0, errors.New("empty string")
	}
	resRegA, errRegA := regexp.MatchString(`^(\*|\d+(\.\d+)*(\.\*)?)$`, a)
	resRegB, errRegB := regexp.MatchString(`^(\*|\d+(\.\d+)*(\.\*)?)$`, b)
	if errRegA != nil || errRegB != nil {
		return 0, errors.New("error in input validation")
	}
	if !resRegA || !resRegB {
		return 0, errors.New("invalid input version")
	}
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
	return result, nil
}
