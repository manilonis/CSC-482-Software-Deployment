#FROM python:2.7
FROM golang:latest


WORKDIR $GOPATH/src/CSC-482-Software-Deployment

COPY . .

# RUN git clone https://github.com/harrytruong/gtfs_realtime_json
ARG LOGGLY_API_KEY
ENV LOGGLY_API_KEY=$LOGGLY_API_KEY
RUN go get -d -v ./...
RUN go install -v ./...
RUN apt-get update && apt-get install -y \
python \
python-pip
#RUN apt-get install python-pip
RUN pip install -r ./gtfs_realtime_json/requirements.txt


CMD ["CSC-482-Software-Deployment"]