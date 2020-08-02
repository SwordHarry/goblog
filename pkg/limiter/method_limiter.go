package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 对接口进行流量限流
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimitInterface {
	return MethodLimiter{Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}}
}

func (m MethodLimiter) Key(c *gin.Context) string {
	panic("implement me")
}

func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	panic("implement me")
}

func (m MethodLimiter) AddBuckets(rules ...BucketRule) LimitInterface {
	panic("implement me")
}
