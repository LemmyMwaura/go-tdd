package main

import (
	// di "github.com/lemmyMwaura/learnTDD/dependancy-injection"
	mock "github.com/lemmyMwaura/learnTDD/mocking"
	// "log"
	// "net/http"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
	mock.Mockmain()
}
