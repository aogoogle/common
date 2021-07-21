package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type JRedis struct {
	*redis.Client
}

func InitRedis(addr, password string, db int) *JRedis {
	jredis := &JRedis{redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})}
	fmt.Println(jredis)
	return jredis
}

//GetKey
//@description: 从redis取key
//@param: key string
//@return: err error, redisJWT string
func (jredis *JRedis)GetKey(key string) (err error, redisJWT string) {
	redisJWT, err = jredis.Client.Get(key).Result()
	return err, redisJWT
}

//SetKey
//@description: key存入redis并设置过期时间
//@param: value string
//@return: err error, redis string
func (jredis *JRedis)SetKey(key string, value string, expiresTime int64) (err error) {
	timer := time.Duration(expiresTime) * time.Second
	err = jredis.Client.Set(key, value, timer).Err()
	return err
}

//GetKeys
//@description: 从redis取key
//@param: userName string
//@return: err error, redisJWT string
func (jredis *JRedis)GetKeys(key string) (err error, items []string) {
	items, err = jredis.Client.Keys(key).Result()
	return err, items
}

// Exists
// @Description: 检测key是否存在
// @Date 2021-05-08 10:22:54
// @receiver jredis
// @param key
// @return bool
func (jredis *JRedis)Exists(key string) bool {
	return jredis.Exists(key)
}

// Delete
// @Description: 删除key
// @Date 2021-05-08 10:23:04
// @receiver JRedis
// @param key
// @return bool
func (JRedis *JRedis)Delete(key string) bool {
	return JRedis.Delete(key)
}