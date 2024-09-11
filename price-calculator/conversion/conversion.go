package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for lineIndex, string := range strings {
		floatPrice, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats[lineIndex] = floatPrice
	}
	return floats, nil
}
