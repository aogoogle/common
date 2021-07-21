package jtime

import (
	"fmt"
	"time"
)

func IsJanuaryMonth() bool {
	return time.Now().Month() == 0
}

func GetCurrentTime() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/1000/1000)
}