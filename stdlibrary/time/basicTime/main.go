package main

import (
	"fmt"
	"time"
)

// time 包学习

// 获取时间对象
func basicTime() {

	now := time.Now() //获取当前时间对象
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)

}

// timezoneDemo 时区示例  Go 语言中使用 location 来映射具体的时区。
func timeZoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds()) //时区偏移量
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC) //相当于创建一个北京时区
	fmt.Println(beijing)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区

	newYork, err := time.LoadLocation("America/New_York") //UTC : 5:00
	if err != nil {
		fmt.Println("加载纽约时区失败", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	/*shanghai, _ := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	tokyo, _ := time.LoadLocation("Asia/Tokyo") // UTC+09:00*/

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	//fmt.Println(timeInUTC, sameTimeInBeijing, sameTimeInNewYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

// Unix.Time
func timeUnix() {
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)

	//时间戳转换为时间对象
	// 获取北京时间所在的东八区时区对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 北京时间 2022-02-22 22:22:22.000000022 +0800 CST
	t := time.Date(2022, 02, 22, 22, 22, 22, 22, beijing)

	var (
		sec  = t.Unix()
		msec = t.UnixMilli()
		usec = t.UnixMicro()
	)
	// 将秒级时间戳转为时间对象（第二个参数为不足1秒的纳秒数）
	timeObj := time.Unix(sec, 22)
	fmt.Println(timeObj)           // 2022-02-22 22:22:22.000000022 +0800 CST
	timeObj = time.UnixMilli(msec) // 毫秒级时间戳转为时间对象
	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
	timeObj = time.UnixMicro(usec) // 微秒级时间戳转为时间对象
	fmt.Println(timeObj)           // 2022-02-22 22:22:22 +0800 CST
}

// 时间间隔-- 运算
func timesOperation() {
	//add 函数
	now := time.Now()
	later := now.Add(time.Hour * 10) // 当前时间添加10小时
	fmt.Println(now, later)

	//sub 函数,求两时间的差值
	a := later.Sub(now)
	fmt.Println(a)

	//equal 比较两个时间是否相同，会考虑时间的因素
	newYork, _ := time.LoadLocation("America/New_York")
	d1 := time.Date(2009, 1, 1, 16, 0, 0, 0, time.UTC)
	d2 := time.Date(2009, 1, 1, 11, 0, 0, 0, newYork)
	b := d1.Equal(d2)
	fmt.Println(b)

	//before 如果t代表的时间点在u之前，返回真；否则返回假。
	fmt.Println(now.Before(later))
	//after 如果t代表的时间点在u之后，返回真；否则返回假。
	fmt.Println(later.After(now))
}

func main() {
	//获取时间，精细到时分秒
	basicTime()
	//时区相关的操作
	timeZoneDemo()
	// Unix时间戳 操作
	timeUnix()
	// 时间运算操作
	timesOperation()
}
