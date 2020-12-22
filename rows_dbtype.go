package odbc

import "github.com/alexbrainman/odbc/api"

var typeMap = map[api.SQLSMALLINT]string{
	api.SQL_UNKNOWN_TYPE: "SQL_UNKNOWN_TYPE",
	api.SQL_CHAR:         "SQL_CHAR",
	api.SQL_NUMERIC:      "SQL_NUMERIC",
	api.SQL_DECIMAL:      "SQL_DECIMAL",
	api.SQL_INTEGER:      "SQL_INTEGER",
	api.SQL_SMALLINT:     "SQL_SMALLINT",
	api.SQL_FLOAT:        "SQL_FLOAT",
	api.SQL_REAL:         "SQL_REAL",
	api.SQL_DOUBLE:       "SQL_DOUBLE",
	api.SQL_DATETIME:     "SQL_DATETIME",
	//api.SQL_DATE:            "SQL_DATE",
	api.SQL_TIME:            "SQL_TIME",
	api.SQL_VARCHAR:         "SQL_VARCHAR",
	api.SQL_TYPE_DATE:       "SQL_TYPE_DATE",
	api.SQL_TYPE_TIME:       "SQL_TYPE_TIME",
	api.SQL_TYPE_TIMESTAMP:  "SQL_TYPE_TIMESTAMP",
	api.SQL_TIMESTAMP:       "SQL_TIMESTAMP",
	api.SQL_LONGVARCHAR:     "SQL_LONGVARCHAR",
	api.SQL_BINARY:          "SQL_BINARY",
	api.SQL_VARBINARY:       "SQL_VARBINARY",
	api.SQL_LONGVARBINARY:   "SQL_LONGVARBINARY",
	api.SQL_BIGINT:          "SQL_BIGINT",
	api.SQL_TINYINT:         "SQL_TINYINT",
	api.SQL_BIT:             "SQL_BIT",
	api.SQL_WCHAR:           "SQL_WCHAR",
	api.SQL_WVARCHAR:        "SQL_WVARCHAR",
	api.SQL_WLONGVARCHAR:    "SQL_WLONGVARCHAR",
	api.SQL_GUID:            "SQL_GUID",
	api.SQL_SIGNED_OFFSET:   "SQL_SIGNED_OFFSET",
	api.SQL_UNSIGNED_OFFSET: "SQL_UNSIGNED_OFFSET",
	api.SQL_SS_XML:          "SQL_SS_XML",
	api.SQL_SS_TIME2:        "SQL_SS_TIME2",
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
