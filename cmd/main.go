package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	months := map[string]string{
		"января": "Jan",
		"февраля": "Feb",
		"марта": "Mar",
		"апреля": "Apr",
		"мая": "May",
		"июня": "Jun",
		"июля": "Jul",
		"августа": "Aug",
		"сентября": "Sep",
		"октября": "Oct",
		"ноября": "Nov",
		"декабря": "Dec",
	}
	layout := "2 Jan 2006 15:04"
	date := "2 июня 2006 в 12:09"


	month := strings.Split(date, " ")[1]
	date = strings.Replace(date, month, months[month], 1)
	date = strings.Replace(date, " в ", " ", 1)

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}