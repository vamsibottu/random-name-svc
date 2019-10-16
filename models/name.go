package models

// NameSrvResponse is used to hold the response details of the name server
type NameSrvResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}
