package utils

import "strings"

func SanitizeUpdateQuery(query string) string {
	i := strings.LastIndex(query, ",")
	query = query[:i] + strings.Replace(query[i:], ",", "", 1)
	query = strings.Replace(query, "[", "0", 1)
	query = strings.Replace(query, "]", "0", 1)

	return query
}

func SanitizeSelectQuery(query string, initLen int) string {
	if len(query) == initLen {
		query = strings.Replace(query, "WHERE", "", 1)
		return query
	}

	i := strings.LastIndex(query, "AND")
	query = query[:i] + strings.Replace(query[i:], "AND", "", 1)
	return query
}
