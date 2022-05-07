package internal

type Data map[string]string
type Items []Data

func (data Data) Data(key string) string {
	return data[key]
}

type Result struct {
	Data Data
	Items Items
}

func (result Result) GetData(key string) string {
	return result.Data.Data(key)
}

func (result Result) GetItems(index int) Data {
	if index >= len(result.Items) || index < 0 {
		return Data{}
	}
	return result.Items[index]
}