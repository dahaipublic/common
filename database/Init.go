package database

import (
	//. "common"
	//	. "github.com/dahaipublic/common/conf"

	//xmysql "github.com/dahaipublic/common/database/mysql"
	xredis "github.com/dahaipublic/common/database/redis"

	"gorm.io/gorm"
	//	"github.com/dahaipublic/common/model"
	//"fmt"
)

var (
	ORMDB *gorm.DB       = nil
	Redis *xredis.CRedis = nil
)

// func InitDB() {
// 	mysql := &Conf.Mysql
// 	redis := &Conf.Redis
// 	//sql链接
// 	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Name, mysql.Charset, true, "Local")

// 	DB = xmysql.GormNew(dataSourceName, mysql.MaxIdleConn, mysql.MaxOpenConn)
// 	r := xredis.New(redis.Host, redis.PassWord, redis.Db)
// 	Redis = xredis.NewRedis(r)

// 	// 初始化创建数据表
// 	model.InitTable(DB)

// 	NewIDWorker(Redis.GetWorkerID())
// 	return
// }
