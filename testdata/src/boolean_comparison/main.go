package main

import (
	"github.com/doug-martin/goqu/v9"
)

func main() {
	goqu.C("bool_column").Eq(true)
}
