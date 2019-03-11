package main

import(
	"fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main(){
sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-2")},
)

svc := dynamodb.New(sess)

result, err := svc.ListTables(&dynamodb.ListTablesInput{})

if err != nil {
    fmt.Println(err)
    os.Exit(1)
}

fmt.Println("Tables:")
fmt.Println("")

for _, n := range result.TableNames {
    fmt.Println(*n)
}

}

