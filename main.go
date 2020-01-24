package main

import ld "gopkg.in/launchdarkly/go-server-sdk.v4"
import "fmt"
import "time"

func main() {
	ldClient, err := ld.MakeClient("SDK-KEY", 5*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}

	key := "user@test.com"

	showFeature, err := ldClient.BoolVariation("show-widgets", ld.NewUser(key), false)
	if err != nil {
		fmt.Println(err.Error())
	}

	if showFeature {
		println("feature is showing")
	  // application code to show the feature
	} else {
		println("feature is NOT showing")
	  // the code to run if the feature is off
	}
}
