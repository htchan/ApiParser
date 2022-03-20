package ApiParser

import (
	"regexp"
	"path/filepath"
	"os"
	"gopkg.in/yaml.v2"
	"strings"
	"errors"

	// "fmt"
)

var baseDirectory string

type Format struct {
	Data []string `yaml:"data"`
	Items string `yaml:"items"`
}

type Result struct {
	Data map[string]string
	Items []map[string]string
}

func Setup(directory string) {
	baseDirectory = directory
}

func identifier2Format(identifier string) (f *Format, err error) {
	defer Recover(func() {
		f = nil
		if err == nil { err = errors.New("unknown identifier") }
	})
	data := strings.Split(identifier, ".")
	filename := filepath.Join(baseDirectory, data[0] + ".yaml")
	key := data[1]
	fileBytes, err := os.ReadFile(filename)
	CheckError(err)
	var fileContent map[string]*Format
	yaml.Unmarshal(fileBytes, &fileContent)
	result, ok := fileContent[key]
	if !ok {
		return nil, errors.New("unknown identifier")
	}
	return result, nil
}

func parseResponseWithFormat(response string, format Format) *Result {
	result := new(Result)
	// find data
	result.Data = make(map[string]string)
	for _, dataRegex := range format.Data {
		regex, err := regexp.Compile(dataRegex)
		if err != nil { continue }
		match := regex.FindStringSubmatch(response)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" && i < len(match) {
				result.Data[name] = match[i]
			}
		}
	}
	
	// find items
	itemsRegex, err := regexp.Compile(format.Items)
	if err != nil { return result }
	matches := itemsRegex.FindAllStringSubmatch(response, -1)
	result.Items = make([]map[string]string, len(matches))
	for i, match := range matches {
		result.Items[i] = make(map[string]string)
		for j, name := range itemsRegex.SubexpNames() {
			if j != 0 && name != "" {
				result.Items[i][name] = match[j]
			}
		}
	}

	return result
}

func Parse(response string, identifier string) *Result {
	format, err := identifier2Format(identifier)
	if err != nil { return nil}
	return parseResponseWithFormat(response, *format)
}