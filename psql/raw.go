package psql

import (
	"github.com/allhingeslost/pgcraft-go"
	"github.com/allhingeslost/pgcraft-go/expr"
)

func RawQuery(query string, args ...any) pgcraft.BaseQuery[expr.Clause] {
	return expr.RawQuery(query, args...)
}
