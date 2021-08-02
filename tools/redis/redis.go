package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
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
func (jredis *JRedis)Delete(key string) bool {
	return jredis.Delete(key)
}

// GetKeyForDB
// @Description: 切换到toDb完后会自动切回db
// @receiver jredis
// @param key
// @param db 默认db
// @param toDb 要切到哪个db
// @return interface{}
// @create 2021-08-02 16:15:05
func (jredis *JRedis)GetKeyForDB(key string, db int, toDb int) interface{} {
	pipe := jredis.Pipeline()
	pipe.Do("select", toDb)
	pipe.Get(key).Result()
	pipe.Do("select", db)
	if cmders, err := pipe.Exec(); err == nil && len(cmders) == 3 {
		result := jredis.GetCmdResult(cmders)
		return result[1]
	}
	return nil
}

// GetCmdResult
// @Description: 获取cmder中的结果
// @receiver jredis
// @param cmders
// @return map[int]interface{}
// @create 2021-08-02 15:20:03
func (jredis *JRedis)GetCmdResult(cmders []redis.Cmder) map[int]interface{} {
	strMap := make(map[int]interface{}, len(cmders))
	for idx, cmder := range cmders {
		//*ClusterSlotsCmd 未实现
		switch reflect.TypeOf(cmder).String() {
		case "*redis.Cmd":
			cmd := cmder.(*redis.Cmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringCmd":
			cmd := cmder.(*redis.StringCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.SliceCmd":
			cmd := cmder.(*redis.SliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringSliceCmd":
			cmd := cmder.(*redis.StringSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStringMapCmd":
			cmd := cmder.(*redis.StringStringMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringIntMapCmd":
			cmd := cmder.(*redis.StringIntMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolCmd":
			cmd := cmder.(*redis.BoolCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolSliceCmd":
			cmd := cmder.(*redis.BoolSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.IntCmd":
			cmd := cmder.(*redis.IntCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.FloatCmd":
			cmd := cmder.(*redis.FloatCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StatusCmd":
			cmd := cmder.(*redis.StatusCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.TimeCmd":
			cmd := cmder.(*redis.TimeCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.DurationCmd":
			cmd := cmder.(*redis.DurationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStructMapCmd":
			cmd := cmder.(*redis.StringStructMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XMessageSliceCmd":
			cmd := cmder.(*redis.XMessageSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XStreamSliceCmd":
			cmd := cmder.(*redis.XStreamSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingCmd":
			cmd := cmder.(*redis.XPendingCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingExtCmd":
			cmd := cmder.(*redis.XPendingExtCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZSliceCmd":
			cmd := cmder.(*redis.ZSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZWithKeyCmd":
			cmd := cmder.(*redis.ZWithKeyCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.CommandsInfoCmd":
			cmd := cmder.(*redis.CommandsInfoCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoLocationCmd":
			cmd := cmder.(*redis.GeoLocationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoPosCmd":
			cmd := cmder.(*redis.GeoPosCmd)
			strMap[idx], _ = cmd.Result()
			break
		}
	}
	return strMap
}