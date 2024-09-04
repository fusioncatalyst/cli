package utils

import (
	"encoding/json"
	"strings"
)

func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func StringifyJSON(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func UnescapeJSONString(str string) string {
	unescapedString := strings.Replace(str, `\"`, `"`, -1)
	unescapedString = strings.Replace(unescapedString, `\n`, "", -1)
	unescapedString = strings.Trim(unescapedString, `"`)
	return str
}

func JSONStringToMap(str string) map[string]interface{} {
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(str), &data)
	return data
}
