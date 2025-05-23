package controller

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func CacheKey(r *http.Request) string {
	params := r.URL.Query()
	sortedKeys := make([]string, 0, len(params))
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	var queryStr string
	for _, k := range sortedKeys {
		for _, v := range params[k] {
			queryStr += fmt.Sprintf("%s=%s&", k, v)
		}
	}

	key := fmt.Sprintf("%s?%s", r.URL.Path, strings.TrimRight(queryStr, "&"))

	return key
}
