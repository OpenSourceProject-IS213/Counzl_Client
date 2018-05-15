package converter

import (
	"strconv"
	"time"
)

func UnixTimeConvert(unixTime int64) time.Time {
	timeString := strconv.FormatInt(unixTime, 10)
	i, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}

//Returnerer string med nåværende tidspunkt i riktig tidssone i følgende format: [hh:mm:ss:ms].
func GetTime() string {
	loc, _ := time.LoadLocation("Europe/Oslo")
	tidspunkt := time.Now().In(loc)
	timestamp := tidspunkt.Format("2006.01.02 15:04:05.0")
	return timestamp
}
