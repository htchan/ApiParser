package ApiParser

import (
	"os"
	"io/fs"
	"fmt"
	"strings"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"github.com/htchan/ApiParser/internal"
)

const CONFIG_EXTENSION = ".yaml"

var defaultFormatSet internal.FormatSet

func NewFormatSet(key string, items string, data ...string) internal.FormatSet {
	return internal.FormatSet{
		key: internal.Format{
			Data: data,
			Items: items,
		},
	}
}

func FromFile(location string) internal.FormatSet {
	filename := strings.ReplaceAll(filepath.Base(location), CONFIG_EXTENSION, "")
	var content map[string]*internal.Format
	f, err := os.Open(location)
	defer f.Close()
	if err != nil {
		return nil
	}
	
	err = yaml.NewDecoder(f).Decode(&content)
	if err != nil {
		return nil
	}

	result := make(internal.FormatSet)
	for key, item := range content {
		result[fmt.Sprintf("%v.%v", filename, key)] = *item
	}
	return result
}

func FromDirectory(path string) internal.FormatSet {
	result := make(internal.FormatSet)
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

func SetDefault(sets ...internal.FormatSet) {
	defaultFormatSet = make(internal.FormatSet)
	for _, set := range(sets) {
		for key, format := range set {
			defaultFormatSet[key] = format
		}
	}
}

func Get(key string) internal.Format {
	return defaultFormatSet.Get(key)
}

func Parse(key string, content string) internal.Result {
	return defaultFormatSet.Get(key).Parse(content)
}