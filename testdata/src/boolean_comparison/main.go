package main

import (
	"github.com/doug-martin/goqu/v9"
)

func main() {
	goqu.I("1").Eq(true)           // want `avoid comparing with boolean value, compare with integer`
	goqu.L("bool_column").Eq(true) // want `avoid comparing with boolean value, compare with integer`
	goqu.C("bool_column").Eq(true) // want `avoid comparing with boolean value, compare with integer`

	goqu.I("bool_column").Eq(0)
	goqu.L("bool_column").Eq(0)
	goqu.C("bool_column").Eq(0)

	goqu.C("bool_column").Eq(1)

	var bl bool
	goqu.C("bool_column").Eq(bl) // want `avoid comparing with boolean value, compare with integer`

	goqu.C("bool_column").IsTrue()  // want `avoid using IsTrue\(\) method, compare with integer`
	goqu.C("bool_column").IsFalse() // want `avoid using IsFalse\(\) method, compare with integer`

	goqu.I("bool_column").IsTrue()  // want `avoid using IsTrue\(\) method, compare with integer`
	goqu.I("bool_column").IsFalse() // want `avoid using IsFalse\(\) method, compare with integer`

	goqu.L("bool_column").IsTrue()  // want `avoid using IsTrue\(\) method, compare with integer`
	goqu.L("bool_column").IsFalse() // want `avoid using IsFalse\(\) method, compare with integer`

	i := 0
	goqu.C("bool_column").Eq(i)

	goqu.C("bool_column").Eq(goqu.L("TRUE"))
	goqu.C("bool_column").Eq(goqu.L("FALSE"))

	_ = goqu.Ex{"bool_column": true} // want `avoid comparing with boolean value, compare with integer`
	_ = goqu.Ex{"bool_column": goqu.L("TRUE")}

	_ = goqu.ExOr{"bool_column": true} // want `avoid comparing with boolean value, compare with integer`
	_ = goqu.ExOr{"bool_column": goqu.L("TRUE")}

	_ = goqu.ExOr{
		"bool_column": true, // want `avoid comparing with boolean value, compare with integer`
		"int_column":  1,
	}
}
