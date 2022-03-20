package ApiParser

import (
	"testing"
)

func Test_identifier2Format(t *testing.T) {
	Setup(".")

	t.Run("success", func(t *testing.T) {
		format, err := identifier2Format("test_data.test_key")
		if err != nil || format == nil || len(format.Data) != 1 ||
			format.Items != "{ \"name\": \"(?P<Name>.*?)\" }" {
				t.Fatalf("identifier2Format fail: format: %v, err: %v", format, err)
			}
	})

	t.Run("file not found", func(t *testing.T) {
		format, err := identifier2Format("unknown_test_data.test_key")
		if err == nil || format != nil {
			t.Fatalf("identifier2Format success: format: %v, err: %v", format, err)
		}
	})

	t.Run("key not found", func(t *testing.T) {
		format, err := identifier2Format("test_data.unknown_test_key")
		if err == nil || format != nil {
			t.Fatalf("identifier2Format success: format: %v, err: %v", format, err)
		}
	})
}

func Test_parseResponseWithFormat(t *testing.T) {
	format := Format{
		Data: []string { "\"Page\": (?P<Page>.*?)," },
		Items: "{ \"name\": \"(?P<Name>.*?)\" }",
	}

	t.Run("success", func(t *testing.T) {
		response := "{\n\"Page\": 5,\n\"data\": [\n{ \"name\": \"name 1\" },\n{ \"name\": \"name 2\" },\n{ \"name\": \"name 3\" },\n]\n}"
		result := parseResponseWithFormat(response, format)
		if result == nil || result.Data["Page"] != "5" || len(result.Items) != 3 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})

	t.Run("data not found", func(t *testing.T) {
		response := "{\n\"data\": [\n{ \"name\": \"name 1\" },\n{ \"name\": \"name 2\" },\n{ \"name\": \"name 3\" },\n]\n}"
		result := parseResponseWithFormat(response, format)
		if result == nil || result.Data["Page"] != "" || len(result.Items) != 3 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})

	t.Run("items not found", func(t *testing.T) {
		response := "{\n\"Page\": 5,\n\"data\": [\n{\"name\": \"name 1\" },\n{\"name\": \"name 2\" },\n{\"name\": \"name 3\" },\n]\n}"
		result := parseResponseWithFormat(response, format)
		if result == nil || result.Data["Page"] != "5" || len(result.Items) != 0 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})
}

func Test_Parse(t *testing.T) {
	Setup(".")
	identifier := "test_data.test_key"
	t.Run("success", func(t *testing.T) {
		response := "{\n\"Page\": 5,\n\"data\": [\n{ \"name\": \"name 1\" },\n{ \"name\": \"name 2\" },\n{ \"name\": \"name 3\" },\n]\n}"
		result := Parse(response, identifier)
		if result == nil || result.Data["Page"] != "5" || len(result.Items) != 3 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})

	t.Run("data not found", func(t *testing.T) {
		response := "{\n\"data\": [\n{ \"name\": \"name 1\" },\n{ \"name\": \"name 2\" },\n{ \"name\": \"name 3\" },\n]\n}"
		result := Parse(response, identifier)
		if result == nil || result.Data["Page"] != "" || len(result.Items) != 3 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})

	t.Run("items not found", func(t *testing.T) {
		response := "{\n\"Page\": 5,\n\"data\": [\n{\"name\": \"name 1\" },\n{\"name\": \"name 2\" },\n{\"name\": \"name 3\" },\n]\n}"
		result := Parse(response, identifier)
		if result == nil || result.Data["Page"] != "5" || len(result.Items) != 0 {
			t.Fatalf("parseResponseWithFormat fail - result: %v", result)
		}
	})
}