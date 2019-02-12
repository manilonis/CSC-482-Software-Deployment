package main

import loggly "github.com/jamespearly/loggly"
import "fmt"

var LOGGLY_TOKEN = "53e8e98e-dc5c-48ea-9a03-8531900fef00"

func main() {
	client := loggly.New("MyApplication")

	err := client.EchoSend("debug", "This is a debug message");

	fmt.Println(err)
}