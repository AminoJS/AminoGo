package utils

import (
	"fmt"
	"os"
)

func DebugLog(file_name string, msg interface{}) {
	if os.Getenv("GO_DEBUG") == "true" {
		fmt.Printf("[%s] [DEBUG] %s\n", file_name, msg)
	}
}
