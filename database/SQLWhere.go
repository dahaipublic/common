package database

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/dahaipublic/common"
)

/*
*@note 数据库不定长度的where语句，sql拼接工具
*@ex
	tplQuerySQL := `
		SELECT *
		FROM task WHERE #W
		ORDER BY create_time DESC LIMIT ?,?`
	tplCountSQL := "SELECT COUNT(*) FROM task WHERE #W"

	sqlWhere := NewSQLWhere(8, "?")
	sqlWhere.AppendInt("task_id=?", taskID).AppendStr("task_tpl_id=?", taskTplID)
	if beginTime != 0 {
		sqlWhere.AppendMutiArg("create_time between ? and ?", beginTime, endTime)
	}

	sqlWhere.AppendStr("dm=?", dm).AppendLike("mc like ?", mc)
	if status == "undone" {
		sqlWhere.AppendNoArg("status<>'done'")
	} else {
		sqlWhere.AppendStr("status=?", status)
	}

	sqlWhere.AppendMutiArg("audit_status IN(?,?)", EAuStatusNormal, EAuStatusRejected)
	querySQL, countSQL, queryArgs, countArgs := sqlWhere.BuildQuerySQLEx(
		tplQuerySQL, tplCountSQL, "#W", offset, size)

	db := mydb.GetDB()
	row := db.QueryRow(countSQL, countArgs...)
	info := &TTaskList{}
	if err := row.Scan(&info.Total); err != nil && err != sql.ErrNoRows {
		return nil, Err_Oracle
	}
	rows, err := db.Query(querySQL, queryArgs...)

*/

type CSQLWhere struct {
	where     []string
	args      []interface{}
	placeHold string
}

// placeHold, 不同数据库传不同标志位， mysql用?   oracle用:   pg用$
func NewSQLWhere(size int, placeHold string) *CSQLWhere {
	where := make([]string, 0, size)
	args := make([]interface{}, 0, size)
	return &CSQLWhere{where, args, placeHold}
}

/*
 *@note 通用函数，不做arg特殊情况特殊处理，但是支持所有的go的内建类型
 *@example where="name=?", arg="li"
 *@example where="state=?", arg=1
 *@example where="flag=?", arg=true
 */
func (this *CSQLWhere) Append(where string, arg interface{}) *CSQLWhere {
	this.where = append(this.where, where)
	this.args = append(this.args, arg)
	return this
}

/*
 *@note 用于当arg为go内建类型的指针的场景，如果且为空需要过滤的场景
 *@example where="name=?", arg=nil
 *@example where="state=?", state=*int arg=flag
 *@example where="flag=?", flag=*bool arg=flag
 */
func (this *CSQLWhere) AppendPointer(where string, arg interface{}) *CSQLWhere {
	if arg != nil {
		this.where = append(this.where, where)
		this.args = append(this.args, arg)
	}
	return this
}

/*
 *@note 用于当arg为字符串且可能为空串，当为空串时需要过滤的场景
 *@example where="name=?", arg="li"
 *@example where="name=?", arg=""
 */
func (this *CSQLWhere) AppendStr(where string, arg string) *CSQLWhere {
	if arg != "" {
		this.where = append(this.where, where)
		this.args = append(this.args, arg)
	}

	return this
}

/*
 *@note 用于当arg为字符串且可能为空串，当为空串时需要过滤,不为空串就需要不错模糊匹配的场景
 *@example where="name like ?", arg="li"
 *@example where="name lile ?", arg=""
 */
func (this *CSQLWhere) AppendLike(where string, arg string) *CSQLWhere {
	if arg != "" {
		this.where = append(this.where, where)
		this.args = append(this.args, fmt.Sprintf("%%%s%%", arg))
	}

	return this
}

/*
 *@note 用于当arg为int且可能为可能为0，当为0时需要过滤的场景
 *@example where="state=?", arg=1
 *@example where="state=?", arg=0
 */
func (this *CSQLWhere) AppendInt(where string, arg int) *CSQLWhere {
	if arg != 0 {
		this.where = append(this.where, where)
		this.args = append(this.args, arg)
	}

	return this
}
func (this *CSQLWhere) AppendInt64(where string, arg int64) *CSQLWhere {
	if arg != 0 {
		this.where = append(this.where, where)
		this.args = append(this.args, arg)
	}

	return this
}

func (this *CSQLWhere) AppendTime(where string, arg time.Time) *CSQLWhere {
	if arg.Unix() != 0 && !arg.IsZero() {
		this.where = append(this.where, where)
		this.args = append(this.args, arg)
	}

	return this
}

/*
 *@note 用于前面的语句自己做了判断或者默认写死的where条件
 *@example status<>'done'
 */
func (this *CSQLWhere) AppendNoArg(where string) *CSQLWhere {
	this.where = append(this.where, where)
	return this
}

/*
 *@note 用于这种不定参数的情况
 *@example audit_status IN(?,?)
 *@example apply_time between ? and ?
 */
func (this *CSQLWhere) AppendMutiArg(where string, args ...interface{}) *CSQLWhere {
	this.where = append(this.where, where)
	this.args = append(this.args, args...)

	return this
}

/*
 *@note 用于这种不定参数的情况. 而且是变参的情况
 *@example audit_status IN(?,?)
 */
func (this *CSQLWhere) AppendIn(clName string, args ...interface{}) *CSQLWhere {
	if len(args) == 0 {
		return this
	}

	where := bytes.Buffer{}
	where.WriteString(clName)
	where.WriteString(" IN (?")
	for i := 1; i < len(args); i++ {
		where.WriteString(",?")
	}
	where.WriteString(")")

	this.where = append(this.where, where.String())
	this.args = append(this.args, args...)
	return this
}

// func (this *CSQLWhere) AppendArray(where string, args interface{}) *CSQLWhere {
// 	if args != nil {
// 		this.where = append(this.where, where)
// 		this.args = append(this.args, pq.Array(args))
// 	}

// 	return this
// }

// /*
//  *@note 用于一些where以外的参数输入，比如分页
//  *@example LIMIT ?,?
//  */
// func (this *CSQLWhere) AppendArg(args ...interface{}) *CSQLWhere {
// 	this.args = append(this.args, args...)
// 	return this
// }

/*
 *@note 生成最终的查询和求和的sql语句并返回sql参数
 *@param tplQuerySQL 查询sql模板
 *@param tplCountSQL 统计的sql模板
 *@param wherePlacehold where替换的占位符号，默认#W
 *@param args where以外的参数，如分页参数，分组参数
 */
func (this *CSQLWhere) BuildQuerySQLEx(
	tplQuerySQL, tplCountSQL, wherePlacehold string,
	args ...interface{}) (querySQL, countSQL string, queryArgs, countArgs []interface{}) {

	whereStr := this.getWhereStr()
	tplQuerySQL = this.fixQueryPageSQL(tplQuerySQL)

	// 因为静态代码扫描软件的原因，sql模板数据的替换只能用 strings.Replace
	querySQL = strings.Replace(tplQuerySQL, wherePlacehold, whereStr, 1)
	countSQL = strings.Replace(tplCountSQL, wherePlacehold, whereStr, 1)
	countArgs = this.args
	queryArgs = append(this.args, args...)
	return
}

func (this *CSQLWhere) getWhereStr() string {
	whereStr := "1=1"
	if len(this.where) != 0 {
		whereStr = strings.Join(this.where, " AND ")
	}

	num := strings.Count(whereStr, "?")
	if num != len(this.args) {
		msg := fmt.Sprintf("CSQLWhere.getWhereStr err, num=%d, argsNum=%d", num, len(this.args))
		common.Warning(msg)
		panic(fmt.Errorf(msg))
	}

	if this.placeHold == "?" {
		return whereStr
	}

	for index := 1; index <= num+1; index++ {
		// oracle 将? 替换成 :1 之类
		// pg 将?  替换成 $1 之类
		placeHold := fmt.Sprintf("%s%d", this.placeHold, index)
		whereStr = strings.Replace(whereStr, "?", placeHold, 1)
	}

	return whereStr
}

func (this *CSQLWhere) fixQueryPageSQL(tplQuerySQL string) string {
	if this.placeHold != "?" {
		offset := len(this.args)
		for index := 1; index <= 2; index++ {
			// oracle 将? 替换成 :1 之类
			// pg 将?  替换成 $1 之类
			placeHold := fmt.Sprintf("%s%d", this.placeHold, index+offset)
			tplQuerySQL = strings.Replace(tplQuerySQL, "?", placeHold, 1)
		}
	}

	return tplQuerySQL
}

/*
 *@note 生成最终的查询sql语句并返回sql参数，支持select取列表并取列表数量这种
 *@param tplQuerySQL 查询sql模板
 *@param wherePlacehold where替换的占位符号，默认#W
 *@param offset 查询时候数据库的偏移
 *@param size 最大返回条数
 */
func (this *CSQLWhere) BuildQuerySQL(
	tplQuerySQL, wherePlacehold string,
	offset, size int) (querySQL string, queryArgs []interface{}) {

	// 因为静态代码扫描软件的原因，sql模板数据的替换只能用 strings.Replace
	whereStr := this.getWhereStr()
	tplQuerySQL = this.fixQueryPageSQL(tplQuerySQL)
	querySQL = strings.Replace(tplQuerySQL, wherePlacehold, whereStr, 1)
	queryArgs = append(this.args, offset, size)
	return
}

/*
 *@note 生成最终的查询sql语句并返回sql参数，支持select，update，delete 后面的where
 *@param tplQuerySQL 查询sql模板,
 *@param wherePlacehold where替换的占位符号，默认#W
 */
func (this *CSQLWhere) BuildSQL(
	tplQuerySQL, wherePlacehold string) (querySQL string, queryArgs []interface{}) {

	// 因为静态代码扫描软件的原因，sql模板数据的替换只能用 strings.Replace
	whereStr := this.getWhereStr()
	querySQL = strings.Replace(tplQuerySQL, wherePlacehold, whereStr, 1)
	queryArgs = this.args
	return
}
