package utils

import (
	"regexp"
	"strings"
)

func getNotionUrlRegex(baseUrl string) string {
	newBaseUrl := baseUrl
	if !strings.HasSuffix(baseUrl, "/") {
		newBaseUrl += "/"
	}
	return newBaseUrl + "(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))"
}

func GetNotionPageIdsFromMessage(baseUrl string, content string) []string {
	regex := regexp.MustCompile(getNotionUrlRegex(baseUrl))
	matches := regex.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return []string{}
	}
	var results []string
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		if match[2] != "" {
			results = append(results, match[2])
			continue
		}
		results = append(results, match[1])
	}
	return results
}
