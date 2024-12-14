package main

import (
	"github.com/doug-martin/goqu/v9"
)

func main() {
	goqu.C("bool_column").Eq(true) // want `compare boolean value with int`

	goqu.C("bool_column").Eq(0)
	goqu.C("bool_column").Eq(1)

	var bl bool
	goqu.C("bool_column").Eq(bl) // want `compare boolean value with int`

	i := 0
	goqu.C("bool_column").Eq(i)

	goqu.C("bool_column").Eq(goqu.L("TRUE"))
	goqu.C("bool_column").Eq(goqu.L("FALSE"))
}
