package main

import (
	"fmt"
	"presentation/mypackages"
	"time"
)

const (
    timezoneJakarta = "Asia"

)


func main() {

	var timestamp int64
	var count int64
	count = 100

	var time_start int64

	timestamp = 1656475200 

	timeT := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation("Asia/Kolkata")
    now := timeT.In(loc)
	time_start = timestamp - 3600
	start_timeT := time.Unix(time_start, 0)
    start_now := start_timeT.In(loc)
	fmt.Println("During the period from ", start_now, " to ", now)
		fmt.Println()
	mypackages.RetriveFronElasticDb(timestamp)
	mypackages.Examine(count)
	count++

}
