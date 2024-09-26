package main

import (
	"fmt"
	_ "strings"
)

func main() {
	var records []bool
	calendar := Constructor()
	records = append(records, calendar.Book(10, 20))
	records = append(records, calendar.Book(15, 25))
	records = append(records, calendar.Book(20, 30))
	fmt.Println(records)

}
