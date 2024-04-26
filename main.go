package main

import (
	// di "learnTests/dependancy-injection"
	mock "learnTests/mocking"
	// "log"
	// "net/http"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
	mock.Mockmain()
}
