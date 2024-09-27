package main

import (
	"fmt"
	_ "strings"
)

func main() {
	var records []bool
	calendar := Constructor()
	records = append(records, calendar.Book(10, 20))
<<<<<<< HEAD
	records = append(records, calendar.Book(15, 25))
	records = append(records, calendar.Book(20, 30))
=======
	records = append(records, calendar.Book(50, 60))
	records = append(records, calendar.Book(10, 40))
	records = append(records, calendar.Book(5, 15))
	records = append(records, calendar.Book(5, 10))
	records = append(records, calendar.Book(25, 55))
>>>>>>> 1e7fb06... ajourter58(vendredi 27, septembre):731. my calendar II
	fmt.Println(records)

}
