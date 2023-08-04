package main

import (
	"fmt"
)

func main() {
	// 保存任务列表到JSON文件
	err := loadTasksFromJsonFile("tasks.json")
	if err != nil {
		fmt.Println("从文件中任务列表失败:", err)
	}

	for {
		fmt.Println("\x1b[32m" + "--------------请选择操作:" + "\x1b[0m")
		fmt.Println("--------------1. 显示任务列表")
		fmt.Println("--------------2. 添加任务")
		fmt.Println("--------------3. 标记任务完成")
		fmt.Println("--------------4. 删除任务")
		fmt.Println("--------------5. 查询某一个任务状态")
		fmt.Println("--------------6. 退出")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showTasks()
		case 2:
			var title, Description string
			var Priority int
			fmt.Println("请输入任务标题:")
			fmt.Scanln(&title)
			fmt.Println("请输入任务描述:")
			fmt.Scanln(&Description)
			fmt.Println("请输入任务优先级:")
			fmt.Scanln(&Priority)
			addTask(title, Description, Priority)
		case 3:
			fmt.Println("请输入要标记完成的任务序号:")
			var index int
			fmt.Scanln(&index)
			completeTask(index - 1)
		case 4:
			fmt.Println("请输入要删除的任务序号:")
			var index int
			fmt.Scanln(&index)
			deleteTask(index - 1)
		case 5:
			fmt.Println("请输入要查询的任务序号:")
			var index int
			fmt.Scanln(&index)
			tasks[index-1].showTaskComplete()
		case 6:
			fmt.Println("再见!")
			return
		default:
			fmt.Println("无效的选项，请重新选择")
		}
	}
}
