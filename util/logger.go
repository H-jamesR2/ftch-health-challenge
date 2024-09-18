// util/logger.go

package util

import (
	"log"
	"runtime"
)

// LogError logs an error with the function name and message
func LogError(message string, err error) {
	funcName := getFuncName()
	log.Printf("ERROR: %s: %s: %v", funcName, message, err)
}

// LogInfo logs an informational message with the function name
func LogInfo(message string) {
	funcName := getFuncName()
	log.Printf("INFO: %s: %s", funcName, message)
}

// getFuncName retrieves the name of the calling function
func getFuncName() string {
	pc, _, _, ok := runtime.Caller(2) // Adjust the skip level as needed
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}
