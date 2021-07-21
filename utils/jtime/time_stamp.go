package jtime

import (
	"database/sql/driver"
	"fmt"
	"github.com/aogoogle/common/utils"
	"strconv"
	"time"
)

//TimeStamp
//转成时间戳
type TimeStamp time.Time

func (t TimeStamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", utils.Pick1Of2(time.Time(t).IsZero(), 0, time.Time(t).Unix()*1000))
	return []byte(stamp), nil
}

// UnmarshalJSON
// @Description: 反序列化时调用
// @Date 2021-05-07 19:03:50
// @receiver t
// @param data
// @return err
func (t *TimeStamp) UnmarshalJSON(data []byte) (err error) {
	value, err := strconv.ParseInt(string(data), 10, 64)
	*t = TimeStamp(time.Unix(value/1000, 0))
	return
}

func (t TimeStamp) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// Scan valueof time.Time
func (t *TimeStamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeStamp(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// String
// @Description: 用于 fmt.Println 和后续验证场景
// @Date 2021-05-07 19:04:18
// @receiver t
// @return string
func (t TimeStamp) String() string {
	return time.Time(t).Format(TimeFormat)
}