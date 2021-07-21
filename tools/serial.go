package tools

import (
	"math/rand"
	"strings"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0  // 纯数字
	KC_RAND_KIND_LOWER = 1  // 小写字母
	KC_RAND_KIND_UPPER = 2  // 大写字母
	KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

// genRandId
// @Description: 随机字符串
func genRandId(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}

// GetNumberId
// @Description: 取id
func GetNumberId() string {
	return "SM"+string(genRandId(13, KC_RAND_KIND_NUM))
}

// GetRandomString
// @Description: 生成随机字符串
// @param lenght
// @return string
func GetRandomString(lenght int, seed int64) string{
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()+seed))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return strings.ToUpper(string(result))
}