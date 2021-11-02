package services

import (
	"strconv"
)

func ParseUint(number string) (uint, error) {
	parsed, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), err
}
