package main

import (
	"fmt"
	"io"
	"net/http"
)

// 测试http包的server端

// 处理连接方法
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,world")
}
func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":23450", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
