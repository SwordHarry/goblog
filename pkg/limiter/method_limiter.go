package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
	"sync"
)

// 对接口进行流量限流
type MethodLimiter struct {
	*Limiter
	lock sync.Mutex
}

func NewMethodLimiter() LimitInterface {
	return &MethodLimiter{Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}}
}

func (m *MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI // /auth?xxx=xxx
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m *MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
}

func (m *MethodLimiter) AddBuckets(rules ...BucketRule) LimitInterface {
	for _, rule := range rules {
		if _, ok := m.limiterBuckets[rule.Key]; !ok {
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return m
}
