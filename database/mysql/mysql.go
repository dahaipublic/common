package mysql

import (
	"database/sql"
	"fmt"
	"log"

	xredis "github.com/dahaipublic/common/database/redis"
	"github.com/dahaipublic/common/model"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	// database driver
	_ "github.com/go-sql-driver/mysql"
	gromMsql "gorm.io/driver/mysql"
)

func InitDB() {
	mysql := &Conf.Mysql
	redis := &Conf.Redis
	//sql链接
	dsn := MakeMysqlDsn(
		mysql.Host, mysql.Port, mysql.Username, mysql.Password, mysql.Name, mysql.Charset)
	ORMDB = GormNew(dsn, mysql.MaxIdleConn, mysql.MaxOpenConn)
	r := xredis.New(redis.Host, redis.PassWord, redis.Db)
	Redis = xredis.NewRedis(r)
	db, err := ORMDB.DB()
	if err != nil {
		Error("ORMDB get DB failed, err=%s", err.Error())
		panic("ORMDB get DB failed")
	}
	//DB = &CMySQL{CDatabase{db}}
	DB = &CMySQL{}
	DB.SetDB(db)

	// 初始化创建数据表
	model.InitTable(ORMDB)

	NewIDWorker(Redis.GetWorkerID())
	return
}

// // sqlx
// func New(dataSourceName string, maxOpenConns int, maxIdleConns int) (db *sql.DB) {
// 	var err error
// 	db, err = sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		log.Fatalf("open mysql error(%+v)", err)
// 		return nil
// 	}
// 	db.SetMaxOpenConns(maxOpenConns)
// 	db.SetMaxIdleConns(maxIdleConns)
// 	return
// }

// gorm
func GormNew(dataSourceName string, maxOpenConns int, maxIdleConns int) (gormDb *gorm.DB) {
	var err error

	fmt.Println("dataSourceName2", dataSourceName)
	gormDb, err = gorm.Open(gromMsql.Open(dataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "live", //前缀
			SingularTable: true, //禁用表复数
		},
	})
	if err != nil {
		log.Fatalf("open mysql error(%+v)", err)
		return nil
	}

	//非生成环境打印所有sql
	if Conf.SetupMode != "prod" {
		gormDb.Logger = gormDb.Logger.LogMode(logger.Info)
	}

	db, _ := gormDb.DB()
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.Exec("SET NAMES utf8mb4")
	return
}

type CMySQL struct{ CDatabase }

//var DB CMysqlDB

// func PostgresInit() {
// 	DB.Open(DBHostPg, DBPortPg, DBUserPg, DBPwdPg, DBNamePg)
// }

func MakeMysqlDsn(dbHost, dbPort, dbUser, dbPwd, dbName, charset string) string {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		dbUser, dbPwd, dbHost, dbPort, dbName, charset, true, "Local")
	return dataSourceName
}

// func (this *CMySQL) Open(
// 	dbHost, dbPort, dbUser, dbPwd, dbName, charset string,
// 	maxOpenConns, maxIdleConns int) bool {
// 	dsn := MakeMysqlDsn(dbHost, dbPort, dbUser, dbPwd, dbName, charset)
// 	Info("open MySQL dsn: ***host=%s port=%s user=%s", dbHost, dbPort, dbUser)
// 	return this.Open(dsn, maxOpenConns, maxIdleConns)
// }

func (this *CMySQL) Open(dsn string, maxOpenConns, maxIdleConns int) bool {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		Error("postgres connecting fail: %s\n", err.Error())
		return false
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(0)
	this.SetDB(db)
	return this.IsOpen()
}
