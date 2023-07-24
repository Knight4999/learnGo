package main

import (
	"encoding/json"
	"fmt"
)

// Json序列化
type student struct {
	Name string
	ID   int
}
type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list"`
}

func newStudent(name string, id int) student {
	return student{
		Name: name,
		ID:   id,
	}
}
func main() {
	c1 := class{
		"火箭班",
		make([]student, 0, 20),
	}
	//往c1中添加学生
	for i := 0; i < 10; i++ {
		//添加10个学生
		tmpStu := newStudent(fmt.Sprintf("stu%02d", i), i)
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	//Json序列化：Go语言中数据 -> Json数据
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//Json反序列化: Json序列化字符串 -> Go语言识别的数据
	/*jsonStr := `{"Title":"火箭班","Students":[{"Name":"stu00","ID":0},
				{"Name":"北京爷儿","ID":1},{"Name":"重庆人","ID":2},{"Name":"stu03","ID":3},
				{"Name":"stu04","ID":4},{"Name":"stu05","ID":5},{"Name":"stu06","ID":6},
				{"Name":"stu07","ID":7},{"Name":"stu08","ID":8},{"Name":"stu09","ID":9}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	fmt.Printf("%#v\n", c2)*/
}
