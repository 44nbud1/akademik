package util

import "encoding/json"

func StructToString(data interface{}) string {
	msg, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(msg)
}
