package clients

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/random_name_svc/config"
	"github.com/random_name_svc/models"
)

// newHTTPClient returns new httpClient
func newHTTPClient() *http.Client {
	netTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,
	}

	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: netTransport,
	}
}

// GetRandomName is used to get the random name from api
func GetRandomName() (*models.NameSrvResponse, error) {
	req, err := http.NewRequest(http.MethodGet, config.NameServerURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.cia.v1+json")

	log.Printf("GetRandomName request received in clients: %v", req)

	res, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// store the response in resp object
	var resp models.NameSrvResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	log.Printf("GetRandomName response: %v", resp)
	return &resp, nil
}

// GetRandomJokeByName is used to get the random joke from api
func GetRandomJokeByName(firstName, lastName string) (*models.JokeSrvResponse, error) {
	// encode the url params, since there are special characters in it
	params := url.Values{}
	params.Add("firstName", firstName)
	params.Add("lastName", lastName)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v?%v&limitTo=[nerdy]", config.JokeServerURL, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.cia.v1+json")

	log.Printf("GetRandomJokeByName request received in clients: %v", req)

	res, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// store the response in resp object
	var resp models.JokeSrvResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	log.Printf("GetRandomJokeByName response: %v", resp)
	return &resp, nil
}
