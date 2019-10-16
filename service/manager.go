package service

import (
	"log"

	"github.com/random_name_svc/clients"
)

// GetRandomJokeByName will get the random name and respective joke from the apis provided
func GetRandomJokeByName() (string, error) {
	log.Print("GetRandomJokeByName request received in services")

	// get the random name from api
	nd, err := clients.GetRandomName()
	if err != nil {
		return "", err
	}

	// get the joke with the given name
	jd, err := clients.GetRandomJokeByName(nd.Name, nd.Surname)
	if err != nil {
		return "", err
	}

	log.Printf("retrieved random joke by name: %v", jd.Value.Joke)
	return jd.Value.Joke, nil
}
