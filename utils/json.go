package utils

import "encoding/json"

func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func StringifyJSON(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
