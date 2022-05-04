package internal

import (
	"testing"
)

func TestFormat(t *testing.T) {
	f := Format{
		Data: []string{ `"Page": (?P<Page>.*?),` },
		Items: `{"name": "(?P<Name>.*?)"}`,
	}

	t.Run("Parse", func (t *testing.T) {
		t.Run("success", func (t *testing.T) {
			t.Parallel()
			content := `"{"Page": 5,"data": [{"name": "name 1"},{"name": "name 2"},{"name": "name 3"}]}"`
			result := f.Parse(content)
			if d := result.Data("Page"); d != "5" {
				t.Errorf(`data: want "%v", got "%v"`, "5", d)
			}
			if len(result.items) != 3 {
				t.Errorf(`items length: want %v, got %v`, 3, len(result.items))
			}
			if d := result.Items(2).Data("Name"); d != "name 3" {
				t.Errorf(`items 2: want "%v", got "%v"`, "name 3", d)
			}
		})

		t.Run("failed", func (t *testing.T) {
			t.Parallel()
			content := `"{}"`
			result := f.Parse(content)
			if len(result.data) != 0 {
				t.Errorf(`data length: want "%v", got "%v"`, 0, len(result.data))
			}
			if len(result.items) != 0 {
				t.Errorf(`items length: want %v, got %v`, 0, len(result.items))
			}
		})
	})
}

func TestFormatSet(t *testing.T) {
	s := FormatSet{
		"format":  Format{
			Data: []string{ `"Page": (?P<Page>.*?),` },
			Items: `{"name": "(?P<Name>.*?)"}`,
		},
	}

	t.Run("Get", func (t *testing.T) {
		t.Run("existing key", func (t *testing.T) {
			t.Parallel()
			f := s.Get("format")
			if len(f.Data) != 1 || f.Data[0] != `"Page": (?P<Page>.*?),` ||
				f.Items != `{"name": "(?P<Name>.*?)"}` {
					t.Errorf("got %v", f)
				}
		})

		t.Run("not exist key", func (t *testing.T) {
			t.Parallel()
			f := s.Get("not_Exist")
			if len(f.Data) != 0 || f.Items != "" {
				t.Errorf("got %v", f)
			}
		})
	})
}

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
	set := FromFile("../test_data/test_data.yaml")
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
	set := FromDirectory("../test_data")
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