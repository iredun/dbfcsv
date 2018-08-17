package main

import (
	"encoding/csv"
	"os"
	"flag"
	"fmt"
	"time"
	"unicode/utf8"
	"github.com/LindsayBradford/go-dbf/godbf"
)

func main() {
	t0 := time.Now();
	
	delimiter := flag.String("d", "|", "delimiter used to separate fields")
	headers := flag.Bool("h", false, "display headers")
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Create(path+".csv")
	
	dbfTable, err := godbf.NewFromFile(path, "CP866")
	if err != nil { panic(err) }

	comma, _ := utf8.DecodeRuneInString(*delimiter)
	out := csv.NewWriter(file)
	out.Comma = comma

	if *headers { 
		out.Write(dbfTable.FieldNames())
		out.Flush()
	}

	// Output rows
	for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		row := dbfTable.GetRowAsSlice(i)
		out.Write(row)
		out.Flush()
	}

	t1 := time.Now();
    fmt.Printf("Elapsed time: %v\n", t1.Sub(t0));
}
