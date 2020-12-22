package odbc

import "github.com/alexbrainman/odbc/api"

var typeMap = map[api.SQLSMALLINT]string{
	api.SQL_C_CHAR:           "SQL_C_CHAR",
	api.SQL_C_LONG:           "SQL_C_LONG",
	api.SQL_C_SHORT:          "SQL_C_SHORT",
	api.SQL_C_FLOAT:          "SQL_C_FLOAT",
	api.SQL_C_DOUBLE:         "SQL_C_DOUBLE",
	api.SQL_C_NUMERIC:        "SQL_C_NUMERIC",
	api.SQL_C_DATE:           "SQL_C_DATE",
	api.SQL_C_TIME:           "SQL_C_TIME",
	api.SQL_C_TYPE_TIMESTAMP: "SQL_C_TYPE_TIMESTAMP",
	api.SQL_C_TIMESTAMP:      "SQL_C_TIMESTAMP",
	api.SQL_C_BINARY:         "SQL_C_BINARY",
	api.SQL_C_BIT:            "SQL_C_BIT",
	api.SQL_C_WCHAR:          "SQL_C_WCHAR",
	api.SQL_C_DEFAULT:        "SQL_C_DEFAULT",
	api.SQL_C_SBIGINT:        "SQL_C_SBIGINT",
	api.SQL_C_UBIGINT:        "SQL_C_UBIGINT",
	api.SQL_C_GUID:           "SQL_C_GUID",
}

//ColumnTypeDatabaseTypeName 附加的
func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {
	switch col := r.os.Cols[index].(type) {
	case *BindableColumn:
		return typeMap[col.SQLType]
	case *NonBindableColumn:
		return typeMap[col.SQLType]
	}
	return ""
}

//ColumnTypeLength 附加的
func (r *Rows) ColumnTypeLength(index int) (length int64, ok bool) {
	switch col := r.os.Cols[index].(type) {
	case *BindableColumn:
		return int64(col.Size), col.IsVariableWidth
	}
	return -1, false

}
