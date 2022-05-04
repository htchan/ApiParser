package internal

import (
	"os"
	"io/fs"
	"fmt"
	"strings"
	"regexp"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

const CONFIG_EXTENSION = ".yaml"

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

var defaultFormatSet FormatSet

func NewFormatSet(key string, items string, data ...string) FormatSet {
	return FormatSet{
		key: Format{
			Data: data,
			Items: items,
		},
	}
}

func FromFile(location string) FormatSet {
	filename := strings.ReplaceAll(filepath.Base(location), CONFIG_EXTENSION, "")
	var content map[string]*Format
	f, err := os.Open(location)
	defer f.Close()
	if err != nil {
		return nil
	}
	
	err = yaml.NewDecoder(f).Decode(&content)
	if err != nil {
		return nil
	}

	result := make(FormatSet)
	for key, item := range content {
		result[fmt.Sprintf("%v.%v", filename, key)] = *item
	}
	return result
}

func FromDirectory(path string) FormatSet {
	result := make(FormatSet)
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), CONFIG_EXTENSION) {
			for key, item := range FromFile(path) {
				result[key] = item
			}
		}
		return nil
	})
	return result
}

func SetDefault(sets ...FormatSet) {
	defaultFormatSet = make(FormatSet)
	for _, set := range(sets) {
		for key, format := range set {
			defaultFormatSet[key] = format
		}
	}
}

func Get(key string) Format {
	return defaultFormatSet.Get(key)
}

func Parse(key string, content string) Result {
	return defaultFormatSet.Get(key).Parse(content)
}