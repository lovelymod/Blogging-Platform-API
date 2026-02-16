package utils

import (
	"encoding/json"
	"fmt"
)

func JsonPrint(data any) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
