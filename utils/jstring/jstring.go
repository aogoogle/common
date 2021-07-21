package jstring

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func IsContainObj(list []interface{}, target interface{}) bool {
	for item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func IsContainStr(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func IsDig(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsAlphaOnly(s string) bool {
	if s == "" {return false}
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

// IntToBytes
// @Description: 整形转换成字节
// @param n
// @return []byte
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt
// @Description:字节转换成整形
// @param b
// @return int
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}