package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// 实现一个cat命令行操作

// 用cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadString('\n') //按照字节读取，遇到换行符停止
		buf = strings.TrimSpace(buf)
		if buf == "q" || buf == "Q" {
			break
		}
		if err == io.EOF {
			//退出之前将已读到的内容输
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse() //解析命令行参数

	if flag.NArg() == 0 { //果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Args()[i])
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
