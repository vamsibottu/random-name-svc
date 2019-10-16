package models

import (
	"encoding/json"
	"log"
)

// Errors is used to hold the error details
type Errors struct {
	Msg       string `json:"error_message"`
	ErrorCode int    `json:"error_code"`
}

// MarshalErrorBinary is used to marshal the errors before sending out
func MarshalErrorBinary(req *Errors) string {
	errMsg, err := json.Marshal(req)
	if err != nil {
		log.Print("failed to marshal the error message")
		return ""
	}
	return string(errMsg)
}
