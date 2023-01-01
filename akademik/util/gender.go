package util

import (
	"fmt"
	"strings"
)

type Gender int

const (
	Female Gender = iota + 1
	Male
)

func (t Gender) String() string {
	return [...]string{
		"F",
		"M",
	}[t-1]
}

func (t Gender) Enum(s string) (Gender, error) {
	m, ok := map[string]Gender{
		"F": Female,
		"M": Male,
	}[strings.ToUpper(s)]
	if !ok {
		return -1, fmt.Errorf("invalid gender")
	}

	return m, nil
}
