package mypackages

import (
	"strconv"
	"time"
)

func PresentTimeStamp ()(timestamp_now int64){
	tNow := time.Now()
	timestamp_now = tNow.Unix()
	return
}

func TimeStampGneerator(timeperiod string, unit string) (startTime string, endTime string) {
	
	var unix_start_time , unix_end_time int64

	if timeperiod == "day"{
		unix_end_time = Present_time_stamp 
		unix_start_time = unix_end_time - 3600
	}
	
	if unit == "M"{
		unix_start_time = unix_start_time * 1000
		unix_end_time = unix_end_time * 1000
	}
	if unit == "U"{
		unix_start_time = unix_start_time * 1000000
		unix_end_time = unix_end_time * 1000000
	}
	s := strconv.FormatInt(unix_start_time, 10)
	e := strconv.FormatInt(unix_end_time, 10)
	return s, e
}


