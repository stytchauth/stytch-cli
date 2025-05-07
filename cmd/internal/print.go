package internal

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintJSON(v any) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
