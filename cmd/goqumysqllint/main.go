package main

import (
	"log"

	"github.com/utgwkk/goqumysqllint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	analyzer, err := goqumysqllint.NewAnalyzer(nil)
	if err != nil {
		log.Fatal(err)
	}

	singlechecker.Main(analyzer)
}
