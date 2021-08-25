package helper

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const Iso8601 = "2006-01-02"
const Yyyymmddhhmmss = "2006/01/02 15:04:05"
const Rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"

func CalcTimeRange(fromDate string, toDate string) int64 {
	fromDateUnix := toUnix(fromDate)
	toDateUnix := toUnix(toDate)
	return toDateUnix - fromDateUnix
}

func toUnix(date string) int64 {
	t, err := time.Parse(Yyyymmddhhmmss, date)
	if err != nil {
		log.Println(err)
	}
	return t.Unix()
}
