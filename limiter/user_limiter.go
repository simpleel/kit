package limiter

import (
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/core/limit"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

// UserRequestLimiterConfig 表示一个用户限流器配置
type UserRequestLimiterConfig struct {
	Rate, Burst int
	Key         string
	Redis       struct {
		Host, Type, Password string
	}
}

// UserRequestLimiter 表示一个用户限流器
type UserRequestLimiter struct {
	rate, burst  int
	userIdentity string
	key          string
	redis        *redis.Redis
}

// NewUserRequestLimiter 创建针对用户的限流器
func NewUserRequestLimiter(rate, burst int, userIdentity, key string, store *redis.Redis) *UserRequestLimiter {
	return &UserRequestLimiter{
		rate:         rate,
		burst:        burst,
		userIdentity: userIdentity,
		key:          key,
		redis:        store,
	}
}

// Handle 捕获客户端发起的 HTTP 请求处理
func (l *UserRequestLimiter) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求中区分用户，每次请求都会重新创建限流器
		key := fmt.Sprintf("%s-%s", l.key, r.Header.Get(l.userIdentity))
		// 创建限流器仅分配内存
		limiter := limit.NewTokenLimiter(l.rate, l.burst, l.redis, key)
		if limiter.Allow() {
			next(w, r)
		}
	}
}
