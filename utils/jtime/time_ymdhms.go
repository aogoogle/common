package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//自定义json输出2006-01-02 15:04:05 日期时间格式

type TimeYmdHms time.Time

const TimeFormat = "2006-01-02 15:04:05"

// MarshalJSON
// @Description: 序列化时调用
// @Date 2021-05-07 19:03:39
// @receiver t
// @return []byte
// @return error
func (t TimeYmdHms) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeFormat))
	return []byte(formatted), nil
}

// UnmarshalJSON
// @Description: 反序列化时调用
// @Date 2021-05-07 19:03:50
// @receiver t
// @param data
// @return err
func (t *TimeYmdHms) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = TimeYmdHms(time.Time{})
		return
	}
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = TimeYmdHms(now)
	return
}

// Value
// @Description: 写入 mysql 时调用
// @Date 2021-05-07 19:04:00
// @receiver t
// @return driver.Value
// @return error
func (t TimeYmdHms) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// Scan
// @Description: 检出 mysql 时调用
// @Date 2021-05-07 19:04:09
// @receiver t
// @param v
// @return error
func (t *TimeYmdHms) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = TimeYmdHms(tTime)
	return nil
}

// String
// @Description: 用于 fmt.Println 和后续验证场景
// @Date 2021-05-07 19:04:18
// @receiver t
// @return string
func (t TimeYmdHms) String() string {
	return time.Time(t).Format(TimeFormat)
}