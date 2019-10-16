package models

// JokeSrvResponse is used to hold the random joke details
type JokeSrvResponse struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}
