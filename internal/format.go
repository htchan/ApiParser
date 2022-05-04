package internal

import (
	"regexp"
)

type Format struct {
	Data []string `yaml:"data"`
	Items string `yaml:"items"`
}

func (format Format) Parse(content string) Result {
	var result Result
	// find data
	result.data = make(Data)
	for _, dataRegex := range format.Data {
		regex, err := regexp.Compile(dataRegex)
		if err != nil { continue }
		match := regex.FindStringSubmatch(content)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" && i < len(match) {
				result.data[name] = match[i]
			}
		}
	}
	
	// find items
	itemsRegex, err := regexp.Compile(format.Items)
	if err != nil { return result }
	matches := itemsRegex.FindAllStringSubmatch(content, -1)
	result.items = make([]Data, len(matches))
	for i, match := range matches {
		result.items[i] = make(Data)
		for j, name := range itemsRegex.SubexpNames() {
			if j != 0 && name != "" {
				result.items[i][name] = match[j]
			}
		}
	}

	return result
}

type FormatSet map[string]Format

func (set FormatSet) Get(key string) Format {
	return set[key]
}