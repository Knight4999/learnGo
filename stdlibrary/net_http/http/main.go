package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//HTTP协议包

func main() {
	//A()

}

func A() {
	//使用net、http包写一个简单发送http请求的client
	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

// 带参数的Get请求
func B() {
	//借助net/url这个标准库
	apiUrl := "http://127.0.0.1:23456/get"
	//Url 解析
	data := url.Values{}
	data.Set("name", "李白")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println("解析url失败:", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

// Post请求
func C() {
	url := "127.0.0.1:23450/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"

	//josn写法
	contentType := "application/json"
	data := `"name":"北京","age":"18"`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
