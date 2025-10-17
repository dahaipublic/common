package mysql

import (
	"bytes"
	"strings"
)

// 过滤xss
func Escape(v *string) {
	strings.ReplaceAll(*v, "'", "")
	strings.ReplaceAll(*v, "\"", "")
	strings.ReplaceAll(*v, "=", "")
}

// sql where 查询条件 ? and ?
func Like(searchText string, column string) string {
	var b bytes.Buffer
	collection := strings.Split(searchText, " ")
	b.WriteString("(")
	for i, v := range collection {
		v = strings.TrimSpace(v)
		Escape(&v)
		if len(v) == 0 {
			continue
		}
		if i > 0 {
			b.WriteString(" OR ")
		} else {
			b.WriteString(" ")
		}
		b.WriteString(column)
		b.WriteString(" LIKE '%")
		b.WriteString(v)
		b.WriteString("%' ")
	}
	b.WriteString(")")
	return b.String()
}
