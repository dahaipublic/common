package database

import (
	"database/sql"

	. "common"
)

// type IDB interface {
// 	Exec(query string, args ...interface{}) (sql.Result, error)
// 	Query(query string, args ...interface{}) (*sql.Rows, error)
// 	QueryRow(query string, args ...interface{}) *sql.Row
// }

func ReplaceSQLLikeStr(str string) string {
	if str == "" {
		return ""
	}

	return "%" + str + "%"
}

func TestDBRef(funcName, id string, countSQLs ...string) EErrCode {
	db := DB.GetDB()
	total := 0
	for _, countSQL := range countSQLs {
		err := db.QueryRow(countSQL, id).Scan(&total)
		if err != nil {
			Warning("%s count query row fail: %s", funcName, err.Error())
			return Err_Oracle
		}
		if total > 0 {
			Warning("%s Ref id: %s total: %d", funcName, id, total)
			return Err_DelRefData
		}
	}

	return No_Error
}

func TestDBRefByMutiParam(funcName, countSQL string, queryArgs ...any) EErrCode {
	total := 0
	err := DB.GetDB().QueryRow(countSQL, queryArgs...).Scan(&total)
	if err != nil {
		Warning("%s count query row fail: %s", funcName, err.Error())
		return Err_Oracle
	}
	if total > 0 {
		Warning("%s Ref id: %v total: %d", funcName, queryArgs, total)
		return Err_DelRefData
	}

	return No_Error
}

/*
*@note 数据库分页查询工具函数
*@author seven
*@ex 参考例子如下

	type TMessageRecord struct {
		RecordID   string `json:"recordID"`
		MsgStatus  bool   `json:"msgStatus"`
		MsgContent string `json:"msgContent"`
		UserID     string `json:"userID"`
	}

	func TestQueryPage() {
		funcName := "TestQueryPage"
		DB.Open("192.168.1.58", "5432", "postgres", "xx", "alpha_zzhl_business_q")
		tplQuerySQL := `select record_id, msg_status, msg_content, user_id
		      FROM msg.record WHERE #W ORDER BY msg_status ASC LIMIT ? OFFSET ? `
		tplCountSQL := `SELECT COUNT(*) FROM msg.record WHERE #W`

		sqlWhere := NewSQLWhere(8, "$")
		sqlWhere.AppendStr("user_id=?", "4sxctzU8Dqj")

		querySQL, countSQL, queryArgs, countArgs := sqlWhere.BuildQuerySQLEx(
			tplQuerySQL, tplCountSQL, "#W", 10, 0)

		scanFunc := func(this *TMessageRecord, rows *sql.Rows) error {
			funcName := "TMessageRecord.Scan"
			err := rows.Scan(&this.RecordID, &this.MsgStatus, &this.MsgContent, &this.UserID)
			if err != nil {
				Warning("[%s] scan err:%s", funcName, err.Error())
				return err
			}
			return nil
		}

		data, errCode := QueryPage[TMessageRecord](
			funcName, querySQL, countSQL, queryArgs, countArgs, scanFunc)
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QueryPage[T any](
	funcName, querySQL, countSQL string,
	queryArgs, countArgs []any,
	scan func(*T, *sql.Rows) error) (*TPageXX[*T], EErrCode) {

	db := DB.GetDB()
	page := &TPageXX[*T]{}
	row := db.QueryRow(countSQL, countArgs...)
	err := row.Scan(&page.Total)
	if err != nil {
		Warning("[%s] count err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}
	rows, err := db.Query(querySQL, queryArgs...)
	if err != nil {
		Warning("[%s] err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}
	defer rows.Close()

	for rows.Next() {
		obj := new(T)
		err = scan(obj, rows)
		if err != nil {
			Warning("[%s] scan err:%s", funcName, err.Error())
			return nil, Err_Oracle
		}
		page.List = append(page.List, obj)
	}
	err = rows.Err()
	if err != nil {
		Warning("%s DB err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}

	return page, No_Error
}

/*
*@note 数据库全量查询工具函数（不分页）
*@author seven
*@ex 参考例子参考 QueryPage
 */
func Query[T any](
	funcName, querySQL string,
	queryArgs []any,
	scan func(*T, *sql.Rows) error) ([]*T, EErrCode) {

	db := DB.GetDB()
	rows, err := db.Query(querySQL, queryArgs...)
	if err != nil {
		Warning("[%s] err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}
	defer rows.Close()

	list := []*T{}
	for rows.Next() {
		obj := new(T)
		err = scan(obj, rows)
		if err != nil {
			Warning("[%s] scan err:%s", funcName, err.Error())
			return nil, Err_Oracle
		}
		list = append(list, obj)
	}
	err = rows.Err()
	if err != nil {
		Warning("%s DB err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}

	return list, No_Error
}

/*
*@note 数据库全量查询工具函数，只scan一个字段的场景
*@author seven
*@ex 参考例子参考

	func QuerySlice() {
		funcName := "QuerySlice"
		querySQL := `select user_id FROM xtpz.user`
		// data 类型为 []string
		data, errCode := QuerySlice[string](funcName, querySQL)
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QuerySlice[T any](
	funcName, querySQL string,
	queryArgs ...any) ([]T, EErrCode) {

	db := DB.GetDB()
	rows, err := db.Query(querySQL, queryArgs...)
	if err != nil {
		Warning("[%s] err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}
	defer rows.Close()

	list := []T{}
	for rows.Next() {
		var obj T
		err := rows.Scan(&obj)
		if err != nil {
			Warning("[%s] scan err:%s", funcName, err.Error())
			return nil, Err_Oracle
		}
		list = append(list, obj)
	}
	err = rows.Err()
	if err != nil {
		Warning("%s DB err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}

	return list, No_Error
}

/*
*@note 数据库全量查询工具函数，只scan一个字段的场景
*@author seven
*@ex 参考例子参考

	func QueryMap() {
		funcName := "QuerySlice"
		querySQL := `select user_id FROM xtpz.user`
		// data 类型为 map[string]string
		data, errCode := QueryMap[string,string](funcName, querySQL)
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QueryMap[TK comparable, TV any](
	funcName, querySQL string,
	queryArgs ...any) (map[TK]TV, EErrCode) {

	db := DB.GetDB()
	rows, err := db.Query(querySQL, queryArgs...)
	if err != nil {
		Warning("[%s] err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}
	defer rows.Close()

	m := map[TK]TV{}
	for rows.Next() {
		var k TK
		var v TV
		err := rows.Scan(&k, &v)
		if err != nil {
			Warning("[%s] scan err:%s", funcName, err.Error())
			return nil, Err_Oracle
		}
		m[k] = v
	}
	err = rows.Err()
	if err != nil {
		Warning("%s DB err:%s", funcName, err.Error())
		return nil, Err_Oracle
	}

	return m, No_Error
}

func idNameScanFunc(this *TIDName, rows *sql.Rows) error {
	return rows.Scan(&this.ID, &this.Name)
}

/*
*@note 数据库全量查询工具函数，只scan两个字段，第一个ID，第二Name的场景，返回[]*TIDName
*@author seven
*@ex 参考例子参考

	func QueryIDName() {
		funcName := "QueryIDName"
		querySQL := `select user_id, show_name FROM xtpz.user`
		// data 类型为 []*TIDName
		data, errCode := QueryIDName(funcName, querySQL)
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QueryIDName(
	funcName, querySQL string, queryArgs ...any) ([]*TIDName, EErrCode) {
	return Query[TIDName](funcName, querySQL, queryArgs, idNameScanFunc)
}

/*
*@note 数据库分页查询工具函数，只scan两个字段，第一个ID，第二Name的场景，返回*TPageXX[TIDName]
*@author seven
*@ex 参考例子参考 QueryIDName
 */
func QueryIDNamePage(
	funcName, querySQL, countSQL string,
	queryArgs, countArgs []any) (*TPageXX[*TIDName], EErrCode) {
	return QueryPage[TIDName](
		funcName, querySQL, countSQL,
		queryArgs, countArgs, idNameScanFunc)
}

/*
*@note 数据库单行查询工具函数，用户指定scan
*@author seven
*@ex 参考例子参考

	type TUser struct {
		UserID   string
		UserName  string
	}
	func QueryRow() {
		funcName := "QueryRow"
		querySQL := `select user_id, show_name FROM xtpz.user`
		// data 类型为 *TUser
		data, errCode := QueryRow(
			funcName, querySQL, []any{},
		 	func(this *TUser, rows *sql.Row) {
				return rows.Scan(&this.UserID, &this.UserName)
		})
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QueryRow[T any](
	funcName, querySQL string,
	queryArgs []any,
	scan func(*T, *sql.Row) error) (v *T, errCode EErrCode) {

	row := DB.GetDB().QueryRow(querySQL, queryArgs...)
	v = new(T)
	err := scan(v, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return v, Err_SQLNoRows
		}

		Warning("%s scan err:%s", funcName, err.Error())
		return v, Err_Oracle
	}

	return v, No_Error
}

/*
*@note 数据库单行查询工具函数，只有一个值的场景
*@author seven
*@ex 参考例子参考

	func QueryValue() {
		funcName := "QueryValue"
		querySQL := `select show_name FROM xtpz.user`
		// data 类型为 string
		data, errCode := QueryValue[string](funcName, querySQL)
		fmt.Printf("data= %v,err=%v", data, errCode)
	}
*/
func QueryValue[T any](
	funcName, querySQL string,
	queryArgs ...any) (v T, errCode EErrCode) {
	err := DB.GetDB().QueryRow(querySQL, queryArgs...).Scan(&v)
	if err != nil {
		if err == sql.ErrNoRows {
			return v, Err_SQLNoRows
		}

		Warning("%s scan err:%s", funcName, err.Error())
		return v, Err_Oracle
	}

	return v, No_Error
}

func CommonDBMutiUpdate(db IDB, funcName, updateSQL string, updateNum int64, args ...any) EErrCode {
	if db == nil {
		db = DB.GetDB()
	}

	res, err := db.Exec(updateSQL, args...)
	if err != nil {
		Warning("%s DB err = %v", funcName, err.Error())
		return Err_Oracle
	}

	if affNum, _ := res.RowsAffected(); affNum != updateNum {
		Warning("%s DB RowsAffected = %d", funcName, affNum)
		return Err_Oracle
	}
	return No_Error
}

func CommonDBUpdate(db IDB, funcName, updateSQL string, args ...any) EErrCode {
	return CommonDBMutiUpdate(db, funcName, updateSQL, 1, args...)
}

/*
*@note 事务工具函数
*@author seven
*@ex 参考例子如下

	func TestWithDBTx() EErrCode {
		const updateStateSQL = `
			UPDATE flow_engine.tpl SET tpl_state=$1, updated_at=$2 WHERE tpl_id=$3`
		now := time.Now().Unix()
		return WithDBTx(funcName, func(tx *sql.Tx) EErrCode {
			now := time.Now().Unix()
			errCode = CommonDBUpdate(tx, funcName, updateStateSQL, 1, now, "tpl_1")
			if errCode != No_Error {
				return errCode
			}

			errCode = CommonDBUpdate(tx, funcName, updateStateSQL, 1, now, "tpl_2")
			if errCode != No_Error {
				return errCode
			}

			return No_Error
		})
	}
*/
func WithDBTx(funcName string, workFunc func(*sql.Tx) EErrCode) EErrCode {
	db := DB.GetDB()
	tx, err := db.Begin()
	if err != nil {
		//logContent = "数据库事务开启失败"
		Warning("%s 数据库事务开启失败", funcName)
		return Err_Oracle
	}
	defer tx.Rollback()

	errCode := workFunc(tx)
	if No_Error != errCode {
		return errCode
	}

	err = tx.Commit()
	if err != nil {
		Warning("%s 数据库事务提交失败. err=%s", funcName, err.Error())
		return Err_Oracle
	}

	return No_Error
}
