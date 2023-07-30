package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 编写代码利用反射实现一个ini文件的解析器程序。

// 定义ini配置的结构体
type Config struct {
	Name    string `ini:"name"`
	Version int    `ini:"version"`
	Author  string `ini:"author"`
}

// ParseINIFile 解析ini文件，并将结果存入指定结构体当中
func ParseINIFile(filename string, config interface{}) error {
	file, err := os.Open(filename) // 打开INI文件
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var section reflect.Value // 当前解析的节（section）

	v := reflect.ValueOf(config).Elem() // 获取配置结构体的Value
	t := v.Type()                       // 获取配置结构体的Type

	for {
		line, err := reader.ReadString('\n') // 逐行读取INI文件内容
		if err == io.EOF {                   // 文件读取到末尾，结束解析
			break
		} else if err != nil { // 其他错误，返回错误信息
			return err
		}

		line = strings.TrimSpace(line) // 去除行首行尾空格

		// 忽略空行和注释行
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		// 解析节（section）
		if line[0] == '[' && line[len(line)-1] == ']' {
			sectionName := line[1 : len(line)-1]
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				if field.Tag.Get("ini") == sectionName { // 根据tag值匹配对应的section
					section = v.Field(i) // 保存当前解析的节
					break
				}
			}
			continue
		}

		// 解析键值对
		index := strings.Index(line, "=")
		if index == -1 {
			//return fmt.Errorf("Invalid line: %s", line)
			continue
		}
		key := strings.TrimSpace(line[:index])
		value := strings.TrimSpace(line[index+1:])

		// 利用反射设置结构体字段的值
		if section.IsValid() {
			for i := 0; i < section.Type().NumField(); i++ {
				field := section.Type().Field(i)
				if field.Tag.Get("ini") == key { // 根据tag值匹配对应的key
					switch field.Type.Kind() {
					case reflect.String:
						section.Field(i).SetString(value) // 设置字段字符串值
					case reflect.Int:
						if intValue, err := strconv.Atoi(value); err == nil {
							section.Field(i).SetInt(int64(intValue)) // 设置字段整数值
						}
					}
					break
				}
			}
		}
	}

	return nil
}

func main() {
	var config Config // 定义配置结构体变量

	err := ParseINIFile("E:/Config.ini", &config) // 解析INI文件，将结果存入config结构体中
	if err != nil {
		fmt.Println("Failed to parse INI file:", err)
		return
	}

	// 输出解析结果
	fmt.Printf("Name: %s\nVersion: %d\nAuthor: %s\n", config.Name, config.Version, config.Author)
}
