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
			if d := result.GetData("Page"); d != "5" {
				t.Errorf(`data: want "%v", got "%v"`, "5", d)
			}
			if len(result.Items) != 3 {
				t.Errorf(`items length: want %v, got %v`, 3, len(result.Items))
			}
			if d := result.GetItems(2).Data("Name"); d != "name 3" {
				t.Errorf(`items 2: want "%v", got "%v"`, "name 3", d)
			}
		})

		t.Run("failed", func (t *testing.T) {
			t.Parallel()
			content := `"{}"`
			result := f.Parse(content)
			if len(result.Data) != 0 {
				t.Errorf(`data length: want "%v", got "%v"`, 0, len(result.Data))
			}
			if len(result.Items) != 0 {
				t.Errorf(`items length: want %v, got %v`, 0, len(result.Items))
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