package main

import (
"fmt"
"log"
"net/http"
)

type reportCounter struct {
	counter int
}

func main() {
	var rc reportCounter
	http.Handle("/server/", &rc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (rc *reportCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User requested on /server")
	rc.counter++
	s := fmt.Sprintf("report call count: %v", rc.counter)
	if rc.counter > 3{
		s += fmt.Sprintf("\n Stop going here you are annoying")
	}
	fmt.Fprint(w, s)
}