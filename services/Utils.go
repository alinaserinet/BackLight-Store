package services

import (
	"strconv"
)

func ParseParam(param string) (uint, error) {
	parsed, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(parsed), err
}