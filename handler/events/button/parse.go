package button_events

import (
	"strconv"
	"strings"
)

type like struct{}

var (
	Like like
)

func (_ like) CustomID(value string, name string) (int, bool) {
	if strings.HasPrefix(value, name+"_review_") {
		data, err := strconv.Atoi(value[len(name)+8:])
		if err == nil {
			return data, true
		}
	}
	return 0, false
}
