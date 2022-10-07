package utils

import (
	"encoding/json"
	"fmt"
)

func StructToJson[T any](structDate T) []byte {
	e, err := json.Marshal(structDate)
	if err != nil {
		fmt.Println(err)
		return []byte(err.Error())
	}
	return e
}
