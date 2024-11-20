package handlerUtils

import "strings"

func GetName(tokens []string) (name string) {
	length := len(tokens)
	for _, token := range tokens[1 : length-1] {
		name += (token) + " "
	}
	name = strings.Title(name)
	name = strings.TrimSpace(name)
	return
}

func GetURL(tokens []string) (URL string) {
	length := len(tokens)
	for _, token := range tokens[length-1:] {
		URL += (token) + " "
	}
	URL = strings.TrimSpace(URL)
	return
}
