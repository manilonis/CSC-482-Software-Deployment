FROM golang:latest

WORKDIR /home/michael/go/src/HTTPS_Server

COPY . .

RUN go get -d -v ./...
# RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

# New build phase -- create binary-only image
FROM alpine:latest

# Add support for HTTPS and time zones
RUN apk update && \
    apk upgrade

WORKDIR /root/

# Copy files from previous build container
COPY --from=build /home/michael/go/src/HTTPS_Server ./

RUN env && pwd && find .
CMD ["HTTPS_Server"]