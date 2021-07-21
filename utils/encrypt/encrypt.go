package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// MD5
// @Description: md5加密
// @Date 2021-05-10 22:05:42
// @param data
// @return string
func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}

// HmacSha1
// @Description: 使用HMAC-SHA1签名方法对data进行签名
// @Date 2021-05-10 22:04:41
// @param value 消息内容
// @param key 密钥
// @return string
func HmacSha1(value, key string) string{

	bkey := []byte(key)
	mac := hmac.New(sha1.New, bkey)
	mac.Write([]byte(value))
	//进行base64编码
	res := hex.EncodeToString(mac.Sum(nil))
	//res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	//res1 := hex.EncodeToString(mac.Sum(nil))
	//fmt.Println(string(res1), aa)
	return res
}