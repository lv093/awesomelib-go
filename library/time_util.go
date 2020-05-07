package library


/**
根据输入的年月日，计算当天是周几
最著名的计算公式是蔡勒（Zeller）公式。即w=y+[y/4]+[c/4]-2c+[26(m+1)/10]+d-1
 output: w: 0 - 周日;
            1 - 周一;
            2 - 周二;
            3 - 周三;
            4 - 周四;
            5 - 周五;
            6 - 周六;
 公式中的符号含义如下，w：星期；
                  c：世纪-1；
                  y：年(两位数);
                  m：月(m大于等于3，小于等于14，即在蔡勒公式中，某年的1、2月要看作上一年的13、14月来计算，比如2003年1月1日要看作2002年的13月1日来计算)
                  d：日；
                  []代表取整，即只要整数部分
 注：C是世纪数减一，y是年份后两位，M是月份，d是日数。1月和2月要按上一年的13月和14月来算，这时C和y均按上一年取值
 */
func ZellerFunction2Week(year, month, day uint16) string {
	var weekday = [7]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}

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
