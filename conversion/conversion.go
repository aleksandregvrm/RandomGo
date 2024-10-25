package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringsToFloats(strings []string) ([]float64, error) {
	listOfFloats := []float64{}
	for _, string := range strings {
		myFloat, err := strconv.ParseFloat(string, 64)
		if err != nil {
			println(err)
			return nil, errors.New("failed to convert string to float")
		}
		listOfFloats = append(listOfFloats, myFloat)
	}
	return listOfFloats, nil
}
