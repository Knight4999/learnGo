package main

import (
	"errors"
	"fmt"
	"os"
)

// fmt 包

func main() {
	// print 、printf 、 println
	fmt.Print("深圳")
	fmt.Printf("%s %s\n", "上海", "土著") //格式化输出
	fmt.Println("无限制")
	fmt.Println("----------------------------------------")
	//Fprint、Fprintf、Fprintln 系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	fmt.Fprint(os.Stdout, "路飞5档觉醒")
	w, err := os.OpenFile("E:/abc.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		return
	}
	defer w.Close()
	//fmt.Printf("%T\n", w)
	_, err = fmt.Fprintf(w, "%s\n", "生命诚可贵，爱情价更高，若为自由故，两者皆可抛")
	if err != nil {
		fmt.Println("写入失败")
		return
	}
	fmt.Println("写入成功")
	fmt.Fprintln(w, "将军如此")

	fmt.Println("------------------------------------------")
	//Sprint、Sprintf、Sprintln 系列函数会把传入的数据生成并返回一个字符串。
	s1 := fmt.Sprint(123445812783)
	fmt.Printf("%T\n", s1)
	name := "中华人民共和国解放军"

	age := 80
	s2 := fmt.Sprintf("name: %s,age: %d", name, age)
	fmt.Println(s2)

	s3 := fmt.Sprintln("world,世界,砸瓦鲁多")
	fmt.Println(s3)

	fmt.Println("------------------------------------------")
	// Errorf 函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	//err := fmt.Errorf("这个是一个奇怪的错误")
	e := errors.New("毛主席说：实践才能出真知")
	e2 := fmt.Errorf("我听:%w", e) //添加%w参数，使err嵌套
	fmt.Println(e2)
	fmt.Println(errors.Unwrap(e2))
	fmt.Println("------------------------------------------")

	//Scan、Scanf、Scanln 可以在程序运行过程中从标准输入获取用户的输入。
	var (
		id      int
		email   string
		married bool
	)
	//fmt.Scan(&id, &email, &married) //读取按照指定格式输入的内容
	//fmt.Scanf("1:%d,2:%s,3:%t\n", &id, &email, &married)
	fmt.Scanln(&id, &email, &married)
	fmt.Printf("扫描结果 id:%d email:%s married:%t \n", id, email, married)

	//Fscan 系列、

	//Sscan 系列、
}
