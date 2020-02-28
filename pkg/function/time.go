package function

import (
	"time"
	log "github.com/sirupsen/logrus"
)

const timeLayout = "2006/01/02 15:04:05"

func GetTimeLayout() string {
	return timeLayout
}

func CalcTimeRange(fromDate string, toDate string) int64 {
	fromDateUnix := toUnix(fromDate)
	toDateUnix := toUnix(toDate)
	return toDateUnix - fromDateUnix
}

func toUnix(date string) int64 {
	t, err := time.Parse(timeLayout, date)
    if err != nil {
        log.Println(err)
    }
    return t.Unix()
}