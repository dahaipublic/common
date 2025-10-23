package blademaster

import (
	"context"
	"github.com/dahaipublic/common/conf"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io-go-redis/adapter"
	rtypes "github.com/zishang520/socket.io-go-redis/types"
	"github.com/zishang520/socket.io/v2/socket"
)

type CService struct {
	SocketIo *socket.Server
}

var (
	once     sync.Once
	instance *CService
)

func NewService() *CService {
	once.Do(func() {
		opts := socket.DefaultServerOptions()
		opts.SetAllowEIO3(true)
		opts.SetCors(&types.Cors{
			Origin:      "*",
			Credentials: true,
			Headers: []*types.Kv{
				{Key: "Upgrade", Value: "websocket"}, {Key: "Connection", Value: "Upgrade"},
			},
		})
		opts.SetTransports(types.NewSet("polling", "websocket", "webtransport"))

		// Redis 适配器
		redisClient := rtypes.NewRedisClient(context.Background(), redis.NewClient(&redis.Options{
			Addr:     conf.Conf.Redis.Host,
			Username: "",
			Password: conf.Conf.Redis.PassWord,
			DB:       conf.Conf.Redis.Db,
		}))

		opts.SetAdapter(&adapter.RedisAdapterBuilder{
			Redis: redisClient,
			Opts:  &adapter.RedisAdapterOptions{},
		})

		httpServer := types.NewWebServer(nil)
		instance = &CService{
			SocketIo: socket.NewServer(httpServer, opts),
		}
	})
	return instance
}
