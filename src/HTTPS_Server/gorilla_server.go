package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"strconv"
)

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/maniloni/status", StatusHandler)
    r.HandleFunc("/maniloni/all", ListHandler)
    r.HandleFunc("/maniloni/test", Tester)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func Tester(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Server Runnign")
}

func ListHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-1")},)
    svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
    TableName: aws.String("subway"),
	}

	result, err := svc.Scan(input)
	if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case dynamodb.ErrCodeResourceNotFoundException:
            fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
        case dynamodb.ErrCodeInternalServerError:
            fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }
    	return
	}
	fmt.Fprint(w, result)
}


func StatusHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-1")},)
    svc := dynamodb.New(sess)
	input := &dynamodb.DescribeTableInput{
    TableName: aws.String("subway"),
	}

	result, err := svc.DescribeTable(input)
	if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case dynamodb.ErrCodeResourceNotFoundException:
            fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
        case dynamodb.ErrCodeInternalServerError:
            fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }
    	return
	}

	data := "{\"table\": \"subway\", \"ItemCount\":" + strconv.FormatInt(*(result.Table.ItemCount), 10) + "}"
    fmt.Fprint(w, data)
	
}