// 这是一个应用包，我们可以通过 import 在其他调用
package tools

//当前程序调用的包，可以是在 $GOPATH/pkg 目录下。GOPATH允许多个目录，如果是多级目录，就在import里面引入多级目录
import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// 定义一个接口
type RedisClient interface {
	//使用 redis 的 BoolCmd、IntCmd 等实现的方法
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
}
