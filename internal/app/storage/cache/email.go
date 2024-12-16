package cache

import (
	"c_program_edu-gin/utils/encrypt"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type EmailStorage struct {
	redis *redis.Client
}

func NewEmailStorage(redis *redis.Client) *EmailStorage {
	return &EmailStorage{redis: redis}
}

func (e *EmailStorage) Set(ctx context.Context, channel string, email string, code string, exp time.Duration) error {
	_, err := e.redis.WithContext(ctx).Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Del(e.failRow(channel, email))
		pipe.Set(e.row(channel, email), code, exp)
		return nil
	})

	return err
}

func (e *EmailStorage) Get(ctx context.Context, channel string, email string) (string, error) {
	return e.redis.WithContext(ctx).Get(e.row(channel, email)).Result()
}

func (e *EmailStorage) Del(ctx context.Context, channel string, email string) error {
	return e.redis.WithContext(ctx).Del(e.row(channel, email)).Err()
}

func (e *EmailStorage) Verify(ctx context.Context, channel string, email string, code string) bool {
	value, err := e.Get(ctx, channel, email)
	if err != nil || len(value) == 0 {
		return false
	}

	if value == code {
		return true
	}

	// 三分钟内失败超过五次则删除该校验信息
	num := e.redis.WithContext(ctx).Incr(e.failRow(channel, email)).Val()
	if num >= 5 {
		_, _ = e.redis.WithContext(ctx).Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Del(e.row(channel, email))
			pipe.Del(e.failRow(channel, email))
			return nil
		})
	} else if num == 1 {
		e.redis.WithContext(ctx).Expire(e.failRow(channel, email), 3*time.Minute)
	}

	return false
}

func (e *EmailStorage) row(channel string, email string) string {
	return fmt.Sprintf("love_knot:auth:email_code:%s:%s", channel, encrypt.Md5(email))
}

func (e *EmailStorage) failRow(channel string, email string) string {
	return fmt.Sprintf("love_knot:auth:email_code_fail:%s:%s", channel, encrypt.Md5(email))
}
