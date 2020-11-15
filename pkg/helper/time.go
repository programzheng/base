package helper

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const iso8601 = "2006-01-02"
const timeLayout = "2006/01/02 15:04:05"
const rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"

func GetIso8601() string {
	return iso8601
}

func GetTimeLayout() string {
	return timeLayout
}

func GetRfc2822() string {
	return rfc2822
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