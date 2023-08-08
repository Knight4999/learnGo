package main

import (
	"fmt"
	"time"
)

// 时间格式化与反格式化

func timeFormat() {
	now := time.Now() //获取当前时间对象

	// 格式化的模板为 2006-01-02 15:04:05

	//24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))
	//12小时制
	fmt.Println(now.Format("2006-01-02 15:04:05 PM Mon Jan"))
	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96
	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}

// 解析字符串格式的时间
func parseDemo() {
	t1, err := time.Parse("2006-01-02 15:04:05", "2023-06-02 12:10:10") //按照指定格式解析 字符串
	if err != nil {
		fmt.Println("parse failed", err)
		return
	}
	fmt.Println(t1)
	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err := time.Parse(time.RFC3339, "2023-08-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}

// ParseInLocation
func parseInLocationDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2023/08/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}

func main() {
	// 将时间格式化为给定模板的字符串
	timeFormat()
	fmt.Println()
	//将给定格式的字符串转化为时间对象(不带时区)
	parseDemo()
	fmt.Println()
	//将给定格式的字符串转化为时间对象(带时区)
	parseInLocationDemo()
}
