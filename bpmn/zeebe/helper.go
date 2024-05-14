package zeebe

import (
	"strings"
)

func splitTrim(s, sep string) []string {
	groups := strings.Split(s, sep)
	for i, group := range groups {
		groups[i] = strings.TrimSpace(group)
	}
	return groups
}

func GetProperty(p []TProperty, name string) (string, bool) {
	for _, v := range p {
		if name == v.Name {
			return v.Value, true
		}
	}
	return "", false
}

func GetTaskHeader(p []TTaskHeader, key string) (string, bool) {
	for _, v := range p {
		if key == v.Key {
			return v.Value, true
		}
	}
	return "", false
}
