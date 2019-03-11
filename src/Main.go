package main

import "fmt"
import "os"
//import "reflect"
import "os/exec"
import "encoding/json"
//import "strings"
import "github.com/landonp1203/goUtils/loggly"
import "github.com/robfig/cron"
import "os/signal"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

type JsonData struct {
  Header Header   `json:"header"`
  Entity []Entity `json:"entity"`
}
type ReplacementPeriod struct {
  End int `json:"end"`
}
type TripReplacementPeriod struct {
  RouteID           string            `json:"route_id"`
  ReplacementPeriod ReplacementPeriod `json:"replacement_period"`
}
type NyctFeedHeader struct {
  TripReplacementPeriod []TripReplacementPeriod `json:"trip_replacement_period"`
  NyctSubwayVersion     string                  `json:"nyct_subway_version"`
}
type Header struct {
  GtfsRealtimeVersion string         `json:"gtfs_realtime_version"`
  Timestamp           int            `json:"timestamp"`
  Incrementality      int            `json:"incrementality"`
  NyctFeedHeader      NyctFeedHeader `json:"nyct_feed_header"`
}
type Arrival struct {
  Time int `json:"time"`
}
type NyctStopTimeUpdate struct {
  ActualTrack    string `json:"actual_track"`
  ScheduledTrack string `json:"scheduled_track"`
}
type Departure struct {
  Time int `json:"time"`
}
type StopTimeUpdate struct {
  Arrival              Arrival            `json:"arrival"`
  ScheduleRelationship int                `json:"schedule_relationship"`
  NyctStopTimeUpdate   NyctStopTimeUpdate `json:"nyct_stop_time_update"`
  Departure            Departure          `json:"departure,omitempty"`
  StopID               string             `json:"stop_id"`
}
type NyctTripDescriptor struct {
  Direction  int    `json:"direction"`
  IsAssigned bool   `json:"is_assigned"`
  TrainID    string `json:"train_id"`
}
type Trip struct {
  NyctTripDescriptor NyctTripDescriptor `json:"nyct_trip_descriptor"`
  RouteID            string             `json:"route_id"`
  TripID             string             `json:"trip_id"`
  StartDate          string             `json:"start_date"`
}
type TripUpdate struct {
  StopTimeUpdate []StopTimeUpdate `json:"stop_time_update"`
  Trip           Trip             `json:"trip"`
}
type Vehicle struct {
  Timestamp           int  `json:"timestamp"`
  CurrentStopSequence int  `json:"current_stop_sequence"`
  Trip                Trip `json:"trip"`
  CurrentStatus       int  `json:"current_status"`
}
type Entity struct {
  TripUpdate TripUpdate `json:"trip_update,omitempty"`
  ID         string     `json:"id"`
  Vehicle    Vehicle    `json:"vehicle,omitempty"`
}

func main(){

	cmd := exec.Command("./try.sh")
	out, err := cmd.Output()

	if err != nil {
    println(err.Error())
    return
}
    o := string(out)

    loggly.Info(o)

    f, err := os.Create("stuff.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(o)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

    var store JsonData
    err = json.Unmarshal([]byte(o), &store)


    fmt.Println(store.Entity[0].ID)

    data, err := json.Marshal(store.Entity[0].TripUpdate)
    fmt.Println(data)

  

    c := cron.New()
    c.AddFunc("@every 10m", func(){poll()})
    c.Start()

  sig := make(chan os.Signal)
  signal.Notify(sig, os.Interrupt, os.Kill)
  <-sig

}

func poll(){
  cmd := exec.Command("./try.sh")
  out, err := cmd.Output()

  if err != nil {
    println(err.Error())
    return
}
    o := string(out)
    loggly.Info(o)

}

func intoDB(m []byte){
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-2")},
)

  // Create DynamoDB client
  svc := dynamodb.New(sess)


  av, err := dynamodbattribute.MarshalMap(m)
  fmt.Println(av)
  input := &dynamodb.PutItemInput{
    Item: av,
    TableName: aws.String("Test"),
  }

  _, err = svc.PutItem(input)

  if err != nil {
    fmt.Println("Got error calling PutItem:")
    fmt.Println(err.Error())
    os.Exit(1)
  }

}