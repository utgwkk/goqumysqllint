package main

import "github.com/doug-martin/goqu/v9"

func main() {
	dialect := goqu.Dialect("mysql")
	dialect.
		From("table").
		Where(
			goqu.C("bool_column").Eq(true), // want "use Eq(1) or IsTrue() instead of Eq(true)"
		)
}
