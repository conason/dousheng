package utils

import "time"

var templateTime = "2006-01-02 15:04:05"
var loc = "Asia/Shanghai"

func DateToTimestamp(timeStr string) (int64, error) {
	parse, err := time.Parse(templateTime, timeStr)
	if err != nil {
		return 0, err
	}
	return parse.Unix(), nil
}

func TimestampToDate(stamp int64) (time.Time, error) {
	dateStr := time.Unix(stamp, 0).Format(templateTime)
	date, err := time.Parse(templateTime, dateStr)
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
