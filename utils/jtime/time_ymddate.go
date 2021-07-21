package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type YMDDate struct {
	time.Time
}

func (t YMDDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

func (t YMDDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *YMDDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = YMDDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
