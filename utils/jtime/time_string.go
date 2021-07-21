package jtime

import "time"

//时间和字符串间的一些处理

func GetTodayDate(format string) string {
	timeObj := time.Now()
	return timeObj.Format(format)
}

//CompareDatetime
//@return 0相同，1小于，2大于
func CompareDatetime(date1, date2 string) int {
	var flag int = 0
	if date1 != "" && date2 != "" {
		t1, err := time.ParseInLocation(TimeFormat, date1+" 00:00:00", time.Local)
		t2, err := time.ParseInLocation(TimeFormat, date2+" 23:59:59", time.Local)
		if err == nil {
			if t1.Equal(t2) {
				flag = 0
			} else if t1.Before(t2) {
				flag = 1
			} else {
				flag = 2
			}
		}
	}
	return flag
}

func Str2Datetime(date string) time.Time {
	time, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return time
}

func IsContainNowTime(start, end time.Time) bool {
	date := time.Now()
	return date.After(start) && date.Before(end)
}