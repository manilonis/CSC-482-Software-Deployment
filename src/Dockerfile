FROM golang:latest

WORKDIR $GOPATH/src/SoftwareDeploymentAPI

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["SoftwareDeploymentAPI"]