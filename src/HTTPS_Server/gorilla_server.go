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
	//"encoding/json"
)

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/maniloni/status", StatusHandler)
    r.HandleFunc("/maniloni/all", ListHandler)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
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
    //vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    //fmt.Println(vars)
    //fmt.Fprint(w, "Hello World")

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

	//fmt.Fprint(w, result)
	// fmt.Println(*(result.Table.ItemCount))

 //    bytes, err := json.Marshal(data)
 //    if err != nil {
 //        panic(err)
 //    }

    fmt.Fprint(w, data)
	
}