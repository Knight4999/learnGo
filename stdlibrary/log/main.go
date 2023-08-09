package main

import (
	"log"
	"os"
)

// log 包

// 初始化函数
func init() {
	log.SetFlags(log.Lmicroseconds | log.Ldate | log.Llongfile)
	log.Println("这是一条很普通的日志")
	w, err := os.OpenFile("e:/mylog.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		log.Fatal("open file failed")
	}
	log.SetOutput(w)                                            //设置日子文件输出位置
	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds) //设置日志输出格式
}
func main() {
	/*log.Print("这是一条很普通的日志")
	v := "很普通的"
	log.Printf("这是一条%s日志", v)
	log.Println("无语了")
	log.Fatalln("这是一条会触发fatal的日志") // fatal，再日志输出完毕后，执行os.exit(1)
	log.Panicln("这是一条会触发panic的日志")*/

	//使用flags、setFlags函数 设置输出格式

	log.Println("他一定很爱你")
	log.SetPrefix("[AX]") //设置输出前缀
	log.Println("下沙")

	//New函数，创建一个logger对象
	logger := log.New(os.Stdout, "[WZH]", log.Llongfile|log.LstdFlags)
	logger.Println("无敌是多么寂寞")
}
