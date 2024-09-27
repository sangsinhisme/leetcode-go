package main

import (
	"fmt"
	_ "strings"
)

func main() {
	var records []bool
	calendar := Constructor()
	records = append(records, calendar.Book(10, 20))
	records = append(records, calendar.Book(50, 60))
	records = append(records, calendar.Book(10, 40))
	records = append(records, calendar.Book(5, 15))
	records = append(records, calendar.Book(5, 10))
	records = append(records, calendar.Book(25, 55))
	fmt.Println(records)

}
