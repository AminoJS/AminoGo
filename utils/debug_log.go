package utils

import (
	"fmt"
	"os"
)

func DebugLog(fileName string, msg interface{}) {
	if os.Getenv("GO_DEBUG") == "true" {
		fmt.Printf("[%s] [DEBUG] %s\n", fileName, msg)
	}
}
