package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidWeightFormat = errors.New("invalid weight format")

type Weight float32

func (r Weight) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%.2f kg", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Weight) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidWeightFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "kg" {
		return ErrInvalidWeightFormat
	}
	i, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return ErrInvalidWeightFormat
	}
	*r = Weight(i)
	return nil
}
