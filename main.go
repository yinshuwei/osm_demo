package main

import (
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yinshuwei/osm/v2"
	"go.uber.org/zap"
)

// User 用户Model
type User struct {
	ID         *int64
	Email      *string
	Nickname   *string
	CreateTime *time.Time
}

func stringPoint(t string) *string {
	return &t
}

func timePoint(t time.Time) *time.Time {
	return &t
}

func main() {
	o, err := osm.New("mysql", "root:123456@/test?charset=utf8", osm.Options{})
	if err != nil {
		fmt.Println(err.Error())
	}

	logger, _ := zap.NewDevelopment()

	osm.ConfLogger(&Logger{logger}, &Logger{logger}, true, 0)

	{ //添加
		user := User{
			Email:      stringPoint("test@foxmail.com"),
			Nickname:   nil,
			CreateTime: timePoint(time.Now()),
		}
		id, count, err := o.Insert("INSERT INTO user (email,nickname,create_time) VALUES (#{Email},#{Nickname},#{CreateTime});", user)
		if err != nil {
			logger.Error("insert error", zap.Error(err))
		}
		logger.Info("test insert", zap.Int64("id", id), zap.Int64("count", count))
	}

	{ //查询
		user := User{
			Email: stringPoint("test@foxmail.com"),
		}
		var results []User
		count, err := o.SelectStructs("SELECT id,email,nickname,create_time FROM user WHERE email=#{Email};", user)(&results)
		if err != nil {
			logger.Error("test select", zap.Error(err))
		}
		resultBytes, _ := json.Marshal(results)
		logger.Info("test select", zap.Int64("count", count), zap.ByteString("result", resultBytes))
	}

	{ // 更新
		user := User{
			Email:    stringPoint("test@foxmail.com"),
			Nickname: stringPoint("hello"),
		}
		count, err := o.Update("UPDATE user SET nickname=#{Nickname} WHERE email=#{Email}", user)
		if err != nil {
			logger.Error("update error", zap.Error(err))
		}
		logger.Info("test update", zap.Int64("count", count))
	}

	{ //查询
		user := User{
			Email: stringPoint("test@foxmail.com"),
		}
		var results []User
		count, err := o.SelectStructs("SELECT id,email,nickname,create_time FROM user WHERE email=#{Email};", user)(&results)
		if err != nil {
			logger.Error("test select", zap.Error(err))
		}
		resultBytes, _ := json.Marshal(results)
		logger.Info("test select", zap.Int64("count", count), zap.ByteString("result", resultBytes))
	}

	{ //删除
		user := User{
			Email: stringPoint("test@foxmail.com"),
		}
		count, err := o.Delete("DELETE FROM user WHERE email=#{Email}", user)
		if err != nil {
			logger.Error("test delete", zap.Error(err))
		}
		logger.Info("test delete", zap.Int64("count", count))
	}

	{ //关闭
		err = o.Close()
		if err != nil {
			logger.Error("close", zap.Error(err))
		}
	}
}
