package main

import (
	"fmt"
)

var weekday = [7]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}

func main() {
	var a, b, c uint16 = 2020, 6, 1
	fmt.Printf("%d年%d月%d日是:%s\n", a, b, c, ZellerFunction2Week(a, b, c))
}

func ZellerFunction2Week(year, month, day uint16) string {
	var y, m, c uint16
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}
	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	which_week := int(week)
	return weekday[which_week]
}
