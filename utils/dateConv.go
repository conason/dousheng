package utils

import (
	"fmt"
	"time"
)

var templateTime = "2006-01-02 15:04:05"
var loc = "Asia/Shanghai"

func DateToTimestamp(timeStr string) (int64, error) {
	parse, err := time.Parse(templateTime, timeStr)
	fmt.Println(parse)
	if err != nil {
		return 0, err
	}
	return parse.Unix(), nil
}

func TimestampToDate(stamp int64) (time.Time, error) {
	if stamp > time.Now().Unix() {
		stamp /= 1000
	}
	//fmt.Println(stamp)
	timeU := time.Unix(stamp, 0)
	dateStr := timeU.Format(templateTime)
	location, err := time.LoadLocation(loc)
	if err != nil {
		return time.Time{}, err
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, location)
	//fmt.Println(dateStr)
	//date, err := time.Parse(templateTime, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func GetTimestamp() (int64, error) {
	formatTime := time.Now().Format(templateTime)
	location, err := time.LoadLocation(loc)
	if err != nil {
		return 0, err
	}
	inLocationTime, err := time.ParseInLocation("2006-01-02 15:04:05", formatTime, location)
	if err != nil {
		return 0, err
	}
	return inLocationTime.Unix(), nil
}
