package stringutils

import "strings"

func SplitNonEmpty(str, sep string) []string {
	tokens := strings.Split(str, sep)
	ret := make([]string, 0)
	for _, token := range tokens {
		if len(token) > 0 {
			ret = append(ret, token)
		}
	}
	return ret
}
