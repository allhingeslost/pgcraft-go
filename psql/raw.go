package psql

import (
	"github.com/pixeldump/pgcraft-go"
	"github.com/pixeldump/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
