package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimitInterface interface {
	Key(c *gin.Context) string                      // 获取对应限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool) // 获取令牌桶
	AddBuckets(rules ...BucketRule) LimitInterface  // 新增多个令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 令牌桶 相应属性
type BucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多久时间放 N 个令牌
	Capacity     int64         // 令牌桶容量
	Quantum      int64         // 放置的令牌数量N
}
