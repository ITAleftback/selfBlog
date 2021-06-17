/**
 * @Author: Anpw
 * @Description:
 * @File:  limiter
 * @Version: 1.0.0
 * @Date: 2021/6/16 18:12
 */

package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

/**
 * @Author: Anpw
 * @Description: 限流器的通用接口
 * @Key: 获取对应的限流器的键值对的名称
 * @GetBucket: 获取令牌桶
 * @AddBuckets: 新增多个令牌桶
 */

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

/**
 * @Author: Anpw
 * @Description: 存储令牌桶与键值对名称的映射关系
 */

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

/**
 * @Author: Anpw
 * @Description: 存储令牌桶的一些相应规则属性
 * @Key: 自定义键值对的名称
 * @FillInterval: 间隔多久时间放N个令牌
 * @Capacity: 令牌桶的容量
 * @Quantum: 每次到达间隔时间后所放的具体令牌数量
 */

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
