package main

import (
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"study-new-relic/config"
	"time"
)

var app *newrelic.Application

func init() {
	var err error
	app, err = newrelic.NewApplication(
		newrelic.ConfigAppName("Your Application Name"),
		newrelic.ConfigLicense(config.GetConfig().NewRelic.LicenseKey),
	)

	if err != nil {
		fmt.Println(err)
		panic("newrelic error")
	}
}

func main() {
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", handler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("listen error")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello world")
}
