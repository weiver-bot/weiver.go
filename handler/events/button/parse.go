package button

import (
	"strconv"
	"strings"
)

type s struct{}

var (
	Parse s
)

func (_ s) CustomID(value string, name string) (int, bool) {
	if strings.HasPrefix(value, name+"_review_") {
		data, err := strconv.Atoi(value[len(name)+8:])
		if err == nil {
			return data, true
		}
	}
	return 0, false
}
