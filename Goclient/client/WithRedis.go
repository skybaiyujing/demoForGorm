package client

import (
	"GoClient/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

func InitReids() (*redis.Client, context.Context) {
	var (
		ctx = context.Background()
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // 没有密码可以留空
			DB:       0,
		})
	)
	return rdb, ctx
}
func GetStudentWithRedisCache(sno int, rdb *redis.Client, ctx context.Context, out chan<- model.Student, wg *sync.WaitGroup) {
	defer wg.Done() //

	key := fmt.Sprintf("student:%d", sno)

	// 尝试从 Redis 中读取缓存
	val, err := rdb.Get(ctx, key).Result()
	if err == nil {
		var cachedStudent model.Student
		if err := json.Unmarshal([]byte(val), &cachedStudent); err == nil {
			fmt.Println("来自 Redis 缓存：", sno)
			out <- cachedStudent
			return
		}
	}

	// 如果 Redis 无缓存，则调用后端接口
	resp, err := GetStuRawBySno(sno) // 返回的是 CommonResponse[[]model.Student]
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	if len(resp.Data) == 0 {
		fmt.Println("未找到学号对应的学生：", sno)
		return
	}
	student := resp.Data[0]
	// 存入 Redis（序列化为 JSON）
	jsonBytes, _ := json.Marshal(student)
	rdb.Set(ctx, key, jsonBytes, 10*time.Minute) // 设置 10 分钟缓存

	out <- student
}
