package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var allStudnet map[int64]*student

// newStudent 是student类型的构造函数 (别人调用我，我就能给他一个student类型变量)
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 把所有的学生都打印出来
	for k, v := range allStudnet {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

// 添加学生
func addStudent() {
	// 向 allStudnet 中添加一个新的学生
	// 1， 创建一个新的学生
	// 2. 获取用户输入
	var id int64
	var name string
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	// 1.2 造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	// 2. 追加到allStudent这个map中
	allStudnet[id] = newStu
}

// 删除学生
func deleteStudnet() {
	// 1. 请用户输入要删除的学生的序号
	var (
		deleteID int64
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteID)
	// 2. 去allStudent这个map中根据学号删除对应的键值对
	delete(allStudnet, deleteID)
}

func main() {
	// 初始化 （开辟内存空间）
	allStudnet = make(map[int64]*student, 48)
	// 无限循环 打印菜单
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统!")

		fmt.Println(`
		   1. 查看所有学生
		   2. 新增学生
		   3. 删除学生
		   4. 退出
		`)
		fmt.Print("请输入你要干啥:")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudnet()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("滚~")
		}
	}
}
