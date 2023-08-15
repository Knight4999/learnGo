package main

//泛型内容3，接口相关

// with_value interface,只包含方法体的方法
type animal interface {
	run()
	move()
}

// general interface,接口内不光只有方法，还有类型的话。只能用于泛型中
type Int interface {
	~int | ~string
	sleep()
	eat()
}

func main() {
	//var i Int //一般接口，只能用于类型约束，不得用于变量定义
}
