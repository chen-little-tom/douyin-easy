package redis_cache

import (
	"douyin-easy/config"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn

const (
	ReplyOk = "OK"
)

func init() {
	app := config.GlobalConfig
	address := fmt.Sprintf("%s:%d", app.Redis.Host, app.Redis.Port)
	conn, err := redis.Dial("tcp", address, redis.DialDatabase(int(app.Redis.Db)), redis.DialUsername(app.Redis.Username), redis.DialPassword(app.Redis.Password))
	if err != nil {
		log.Printf("redis conn failure,err: %s", err)
		log.Panic("Application stop")
	}
	Conn = conn
}