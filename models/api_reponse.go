package models

import (
	"encoding/json"
	"log"
)

// Response is used to hold the api response details
type Response struct {
	Joke string `json:"joke"`
}

// MarshalResponseBinary is used to marshal the response before sending out
func MarshalResponseBinary(req *Response) []byte {
	msg, err := json.Marshal(req)
	if err != nil {
		log.Print("failed to marshal the response message")
		return nil
	}
	return msg
}
