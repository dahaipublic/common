package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
)

func New(addr string, pwd string, dbNum int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,   // no password set
		DB:       dbNum, // use default DB
	})
	return
}

type CRedis struct {
	redis *redis.Client
}

func NewRedis(r *redis.Client) *CRedis {
	return &CRedis{r}
}

func (d *CRedis) GetRedis() *redis.Client {
	return d.redis
}

func (d *CRedis) GetWorkerID() uint64 {
	cmd := d.redis.Incr("incr:workerid")
	if cmd.Err() != nil {
		return 0
	}

	workerId, _ := cmd.Uint64()
	workerId = (workerId % 1023) + 1
	//Info("worker id: %d", workerId)
	return workerId
}

func (d *CRedis) TxDel(key string) (err error) {
	tx := d.redis.TxPipeline()
	tx.Del(key)
	_, err = tx.Exec()
	return
}

// 删除缓存key
func (d *CRedis) RedisDel(key string) (r bool) {
	d.redis.Del(key)
	return true
}

// 一次性删除多个key
func (d *CRedis) RedisDelMany(keys []string) bool {
	for _, key := range keys {
		d.RedisDel(key)
	}
	return true
}

// 按照前缀关键词删除KEY
func (d *CRedis) RedisDelKeys(key string) (total int) {
	list := d.redis.Keys(key + "*")
	for _, v := range list.Val() {
		d.redis.Del(v)
	}
	return len(list.Val())
}

// 从redis根据键取字符串值
func (d *CRedis) RedisGet(key string) string {
	res := d.redis.Get(key)
	if res.Err() != nil {
		return ""
	}
	return res.Val()
}
func (d *CRedis) RedisResult(key string) string {
	res := d.redis.Get(key)
	if res.Err() != nil {
		return ""
	}
	return res.Val()
}

// 设置redis Key=>Value
func (d *CRedis) RedisSet(key string, value string, expireAt time.Duration) {
	d.redis.Set(key, value, expireAt)
}

// 没有则设置redis Key=>Value
func (d *CRedis) RedisSetNX(key string, value string, expireAt time.Duration) (b bool) {
	res := d.redis.SetNX(key, value, expireAt)
	return res.Val()
}

// redis +1
func (d *CRedis) RedisInc(key string) int64 {
	res := d.redis.Incr(key)
	if res.Err() != nil {
		return 0
	}
	return res.Val()
}

// redis 增加
func (d *CRedis) RedisIncBy(key string, value int64) int64 {
	res, err := d.redis.IncrBy(key, value).Result()
	if err != nil {
		return 0
	}
	return res
}

// 设置redis Key=>Value
func (d *CRedis) RedisSetComm(key string, value interface{}, expireAt time.Duration) {
	d.redis.Set(key, value, expireAt)
}

// 指定前缀的Key列表
func (d *CRedis) RedisKeys(key string) *redis.StringSliceCmd {
	return d.redis.Keys(key)
}

// 设置有序集合的值
func (d *CRedis) RedisZAdd(key string, score float64, member interface{}) bool {
	Z := redis.Z{score, member}
	result := d.redis.ZAdd(key, &Z)
	if result.Err() != nil {
		return false
	}
	return true
}

// 获取有序集合的会员值
func (d *CRedis) RedisZScore(key, member string) float64 {
	result := d.redis.ZScore(key, member)
	if result.Err() == nil {
		return result.Val()
	}
	return 0
}

// 获取有序集合及分数
func (d *CRedis) RedisZRevRangeWithScores(key string, start, stop int64) (res []redis.Z) {
	res, err := d.redis.ZRevRangeWithScores(key, start, stop).Result()
	if err == nil {
		return
	}
	return
}

// 获取有序集合及分数
func (d *CRedis) RedisZRange(key string, start, stop int64) (res []string) {
	res, err := d.redis.ZRange(key, start, stop).Result()
	if err == nil {
		return
	}
	return
}

// 获取有序集合总成员个数
func (d *CRedis) RedisZCard(key string) int64 {
	result := d.redis.ZCard(key)
	if result.Err() == nil {
		return result.Val()
	}
	return 0
}

// 有序集合自增
func (d *CRedis) RedisZIncrBy(key string, increment float64, member string) float64 {
	val, err := d.redis.ZIncrBy(key, increment, member).Result()
	if err != nil {
		return 0
	}
	return val
}

// 删除有序集合中的某条数据
func (d *CRedis) RedisZRem(key, member string) {
	d.redis.ZRem(key, member)
}

// 主有序集合自增
func (d *CRedis) RedisAnchorZIncrBy(key string, increment float64, member string) {
	d.redis.ZIncrBy(key, increment, member)
}

// 删除有序集合中的某条数据
func (d *CRedis) RedisAnchorZRem(key, member string) {
	d.redis.ZRem(key, member)
}

// 无序集合增加成员
func (d *CRedis) RedisSAdd(key string, members string) {
	d.redis.SAdd(key, members)
}

// 是否无序集合成员
func (d *CRedis) RedisSIsMember(key string, members string) bool {
	r := d.redis.SIsMember(key, members)
	return r.Val()
}

// 是否无序集合所有成员
func (d *CRedis) RedisSMember(key string) []string {
	r := d.redis.SMembers(key)
	return r.Val()
}

// 删除无序集合成员
func (d *CRedis) RedisSRem(key string, members string) {
	d.redis.SRem(key, members)
}

// 获取多个无序集合交集
func (d *CRedis) RedisSInter(keys ...string) (res []string) {
	result := d.redis.SInter(keys...)
	if result.Err() == nil {
		res = result.Val()
		return
	}
	return
}

// 获取多个无序集合差集
func (d *CRedis) RedisSDiff(keys ...string) (res []string) {
	result := d.redis.SDiff(keys...)
	if result.Err() == nil {
		res = result.Val()
		return
	}
	return
}

// 哈希表自增
func (d *CRedis) RedisHIncrBy(key, field string, incr int64) (int64, error) {
	result, err := d.redis.HIncrBy(key, field, incr).Result()
	if err != nil {
		return -1, err
	}
	return result, nil
}

// 哈希表设置值
func (d *CRedis) RedisHSet(key string, values ...interface{}) {
	d.redis.HSet(key, values...)
}

// 哈希表取值
func (d *CRedis) RedisHGet(key, field string) (value string) {
	result := d.redis.HGet(key, field)
	if result.Err() != nil {
		return
	}
	value = result.Val()
	return
}

// 删除哈希表
func (d *CRedis) RedisHDel(key string, fields ...string) {
	d.redis.HDel(key, fields...)
}

// 查询无序集合成员总数
func (d *CRedis) RedisSCard(key string) (total int64) {
	result := d.redis.SCard(key)
	if result.Err() != nil {
		total = 0
	} else {
		total = result.Val()
	}
	return
}

// 缓存锁
func (d *CRedis) RedisGetLock(lockKey string) bool {
	//抢锁
	r := d.redis.Get(lockKey)
	if r.Err() != nil {
		//加锁
		d.redis.Set(lockKey, "1", time.Second*60)
		return true
	}

	lock, _ := strconv.ParseInt(r.Val(), 10, 64)
	if lock > 0 {
		return false
	}

	//加锁
	d.redis.Set(lockKey, "1", time.Second*60)
	return true
}

// 解除缓存锁
func (d *CRedis) RedisUnLock(lockKey string) bool {
	d.redis.Del(lockKey)
	return true
}

// 列表数据
func (d *CRedis) RedisLRange(key string, start, stop int64) *redis.StringSliceCmd {
	return d.redis.LRange(key, start, stop)
}

// 添加列表数据
func (d *CRedis) RedisLPush(key, data string) *redis.IntCmd {
	return d.redis.LPush(key, data)
}

// 设置过期时间
func (d *CRedis) RedisExpire(key string, time time.Duration) *redis.BoolCmd {
	return d.redis.Expire(key, time)
}
