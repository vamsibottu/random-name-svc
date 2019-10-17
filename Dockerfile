FROM golang:1.13.1

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/random-name-svc

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

RUN go install github.com/random-name-svc

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/random-name-svc

EXPOSE 5000
