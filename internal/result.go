package internal

type Data map[string]string

func (data Data) Data(key string) string {
	return data[key]
}

type Result struct {
	data Data
	items []Data
}

func (result Result) Data(key string) string {
	return result.data.Data(key)
}

func (result Result) Items(index int) Data {
	if index >= len(result.items) || index < 0 {
		return Data{}
	}
	return result.items[index]
}