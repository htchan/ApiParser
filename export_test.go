package main

import (
	"testing"
)


func TestNewFormatSet(t *testing.T) {
	set := NewFormatSet("key", "items", "data1", "data2")
	f, ok := set["key"]

	if !ok {
		t.Errorf("format not found: %v", f)
	}
	if len(f.Data) != 2 || f.Data[0] != "data1" || f.Data[1] != "data2" {
		t.Errorf(`data: got %v`, f.Data)
	}
	if f.Items != "items" {
		t.Errorf("items: got %v", f.Items)
	}
}

func TestFromFile(t *testing.T) {
	set := FromFile("./test_data/test_data.yaml")
	f, ok := set["test_data.test_key"]

	if !ok {
		t.Errorf("format not found: %v", f)
	}
	if len(f.Data) != 1 || f.Data[0] != `"Page": (?P<Page>.*?),` {
		t.Errorf(`data: got %v`, f.Data)
	}
	if f.Items != `{ "name": "(?P<Name>.*?)" }` {
		t.Errorf(`items: got "%v"`, f.Items)
	}

	_, ok = set["test_data.test_key2"]
	if !ok {
		t.Errorf("format not found: %v", f)
	}
}

func TestFromDirectory(t *testing.T) {
	set := FromDirectory("./test_data")
	f, ok := set["test_data.test_key"]
	if !ok {
		t.Errorf("format not found: %v", f)
	}
	f, ok = set["test_data.test_key2"]
	if !ok {
		t.Errorf("format not found: %v", f)
	}
}

func TestSetDefault(t *testing.T) {
	SetDefault(NewFormatSet("key", "string"))
	if _, ok := defaultFormatSet["key"]; !ok {
		t.Errorf("it didn't change default")
	}
}