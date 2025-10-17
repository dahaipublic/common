package pg

// import (
// 	_ "github.com/lib/pq"
// )

type CPgDB struct{ CDatabase }

// var DB CPgDB

// func PostgresInit() {
// 	DB.Open(DBHostPg, DBPortPg, DBUserPg, DBPwdPg, DBNamePg)
// }

// func (this *CPgDB) Open(dbHost, dbPort, dbUser, dbPwd, dbName string) bool {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		dbHost, dbPort, dbUser, dbPwd, dbName)
// 	Info("open postgres dsn: ***host=%s port=%s user=%s", dbHost, dbPort, dbUser)
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		Error("postgres connecting fail: %s\n", err.Error())
// 		return false
// 	}

// 	db.SetMaxIdleConns(DBMaxIdel)
// 	db.SetMaxOpenConns(DBMaxOpen)
// 	db.SetConnMaxLifetime(0)
// 	this.SetDB(db)
// 	return this.IsOpen()
// }
