## Random Joke Service 

### Steps to run the service

#### Requires:

* ![Install Golang](https://golang.org/doc/install)

* ![Install Golang CI Lint](https://github.com/golangci/golangci-lint)

* ![Install Go Dep](https://github.com/golang/dep)


#### How to run

##### Webserver

## Build Docker container and tag the container with name
```
docker build -t random-name-svc .
```

## Run the container 
```
docker run --publish 6060:8080 --name new-svc --rm random-name-svc
```

## To see the results
```
localhost:6060
```

##### To check the code Quality Golang CI Lint

```
golangci-lint run
```
