package database

import (
	"database/sql"
	//	"fmt"

	"github.com/dahaipublic/common"
)

var DB IDatabase

type IDatabase interface {
	SetDB(db *sql.DB)
	Open(dsn string, maxOpenConns, maxIdleConns int) bool
	IsOpen() bool
	Close()
	GetDB() *sql.DB
	Count(countSQL string, countArgs ...interface{}) int
}

type CDatabase struct{ db *sql.DB }

func (this *CDatabase) SetDB(db *sql.DB) { this.db = db }

func (this *CDatabase) IsOpen() bool {
	err := this.db.Ping()
	if err != nil {
		common.Error("database unreachable! %s", err.Error())
		return false
	}

	return true
}

func (this *CDatabase) Close()         { this.db.Close() }
func (this *CDatabase) GetDB() *sql.DB { return this.db }

func (this *CDatabase) Count(countSQL string, countArgs ...interface{}) int {
	row := this.db.QueryRow(countSQL, countArgs...)
	c := 0
	if err := row.Scan(&c); err != nil {
		common.Warning("CDatabase.Count err=%s", err.Error())
		return 0
	}

	return c
}

type IDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
