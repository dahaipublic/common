package util

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type LimitInfo struct {
	Now   int64
	Count int
}

var userLimitBuy = make(map[int64]*LimitInfo)

var lockBuy sync.Mutex

func Check(uid int64) bool {
	now := time.Now().UnixMilli()
	v, ok := userLimitBuy[uid]
	if ok {
		if now-v.Now <= 200 && v.Count >= 1 {
			return false
		} else {
			v.Now = now
			v.Count = 1
			return true
		}
	} else {
		var limitInfo LimitInfo
		limitInfo.Now = now
		limitInfo.Count = 1
		lockBuy.Lock()
		userLimitBuy[uid] = &limitInfo
		lockBuy.Unlock()
		return true
	}
}

//// key = ip + path，确保每个接口单独计数
//var visitors = make(map[string]*rate.Limiter)
//var mu sync.Mutex
//
//// 每秒最多 5 次，突发 2 次
//func newLimiter() *rate.Limiter {
//	return rate.NewLimiter(5, 2)
//}
//
//func getLimiter(ip, path string) *rate.Limiter {
//	key := ip + ":" + path
//	mu.Lock()
//	defer mu.Unlock()
//
//	limiter, exists := visitors[key]
//	if !exists {
//		limiter = newLimiter()
//		visitors[key] = limiter
//	}
//	return limiter
//}
//
//func RateLimitMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ip := c.ClientIP()
//		path := c.FullPath() // 获取接口路径，比如 /api/data
//
//		limiter := getLimiter(ip, path)
//		if !limiter.Allow() {
//			c.JSON(http.StatusTooManyRequests, gin.H{
//				"error": "Too Many Requests on " + path,
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}

var (
	visitors     = make(map[string]*rate.Limiter) // ip -> limiter
	mu           sync.Mutex
	bannedIPs    = make(map[string]time.Time) // ip -> 解封时间
	violationCnt = make(map[string]int)       // ip -> 连续违规次数
)

// 配置参数
const (
	limitRPS        = 20              // 每秒 10 次
	burst           = 20              // 瞬时突发 10 次
	banDuration     = 5 * time.Minute // 封禁 5 分钟
	maxViolations   = 3               // 连续 3 次违规才封禁
	cleanupInterval = 1 * time.Minute // 清理周期
)

// 初始化限流器
func newLimiter() *rate.Limiter {
	return rate.NewLimiter(limitRPS, burst)
}

// 获取限流器
func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = newLimiter()
		visitors[ip] = limiter
	}
	return limiter
}

// 检查是否封禁
func isBanned(ip string) bool {
	mu.Lock()
	defer mu.Unlock()
	until, exists := bannedIPs[ip]
	if !exists {
		return false
	}
	if time.Now().After(until) {
		delete(bannedIPs, ip)
		delete(violationCnt, ip)
		return false
	}
	return true
}

// 增加违规次数，必要时封禁
func recordViolation(ip string) {
	mu.Lock()
	defer mu.Unlock()

	violationCnt[ip]++
	if violationCnt[ip] >= maxViolations {
		bannedIPs[ip] = time.Now().Add(banDuration)
		violationCnt[ip] = 0 // 重置计数
	}
}

// 定时清理过期数据
func cleanupTask() {
	ticker := time.NewTicker(cleanupInterval)
	for {
		<-ticker.C
		mu.Lock()
		now := time.Now()
		// 清理过期封禁
		for ip, until := range bannedIPs {
			if now.After(until) {
				delete(bannedIPs, ip)
				delete(violationCnt, ip)
			}
		}
		// 清理长期不用的 limiter
		// （这里简单起见，不追踪最后访问时间，如果需要可扩展）
		mu.Unlock()
	}
}

// Gin 中间件
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// 检查是否封禁
		if isBanned(ip) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Too Many Requests (banned)",
			})
			c.Abort()
			return
		}

		limiter := getLimiter(ip)
		if !limiter.Allow() {
			// 记录违规，可能触发封禁
			recordViolation(ip)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
