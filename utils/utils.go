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
	return "^" + newBaseUrl + "(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))$"
}

func GetNotionPageIdFromMessage(baseUrl string, content string) (bool, string) {
	regex := regexp.MustCompile(getNotionUrlRegex(baseUrl))
	matches := regex.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return false, ""
	}
	if len(matches[0]) < 3 {
		return false, ""
	}
	if matches[0][2] != "" {
		return true, matches[0][2]
	}
	return true, matches[0][1]
}
