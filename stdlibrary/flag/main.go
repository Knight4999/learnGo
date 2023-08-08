package main

import (
	"flag"
	"fmt"
	"time"
)

// flag包，用于解析命令行参数

func main() {
	/*//os.Args 用于获取命令行参数
	if len(os.Args) > 0 { //os.Args 返回一个参数切片
		for index, v := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, v)
		}
	}*/

	/*//定义一个flag ，需要注意的是，此时name、age、married、delay均为对应类型的指针
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	fmt.Println(name, age, married, delay)*/

	//使用flag.TypeVar
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
