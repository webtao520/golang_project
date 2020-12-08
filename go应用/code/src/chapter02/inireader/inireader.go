package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 根据文件名，段名，键名获取ini的值
func getValue(filename, expectSection, expectKey string) string {

	// 打开文件
	file, err := os.Open(filename)

	// 文件找不到，返回空
	if err != nil {
		return ""
	}

	// 在函数结束时，关闭文件
	defer file.Close()

	// 使用读取器读取文件
	reader := bufio.NewReader(file)

	// 当前读取的段的名字
	var sectionName string

	for {

		// 读取文件的一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// 切掉行的左右两边的空白字符
		linestr = strings.TrimSpace(linestr)

		// 忽略空行
		if linestr == "" {
			continue
		}

		// 忽略注释
		if linestr[0] == ';' {
			continue
		}

		// 行首和尾巴分别是方括号的，说明是段标记的起止符
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {

			// 将段名取出
			sectionName = linestr[1 : len(linestr)-1]

			// 这个段是希望读取的
		} else if sectionName == expectSection {

			// 切开等号分割的键值对
			pair := strings.Split(linestr, "=")

			// 保证切开只有1个等号分割的简直情况
			if len(pair) == 2 {

				// 去掉键的多余空白字符
				key := strings.TrimSpace(pair[0])

				// 是期望的键
				if key == expectKey {

					// 返回去掉空白字符的值
					return strings.TrimSpace(pair[1])
				}
			}

		}

	}

	return ""
}

func main() {

	fmt.Println(getValue("example.ini", "remote \"origin\"", "fetch"))

	fmt.Println(getValue("example.ini", "core", "hideDotFiles"))
}
