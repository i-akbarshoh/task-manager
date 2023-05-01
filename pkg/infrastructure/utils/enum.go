package utils

import "fmt"

func IsEnumValue(enum []string, value string) bool {
	fmt.Println(value)
	for _, v := range enum {
		if v == value {
			return true
		}
	}
	return false
}
