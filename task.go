package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

const Filename = "tasks.json"

type Task struct {
	Title       string
	Complete    bool
	Description string
	//优先级 0 1 2 3 4 5
	Priority int
}

var tasks []Task

func loadTasksFromJsonFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return json.Unmarshal(data, &tasks)
}

func addTask(title string, Description string, Priority int) {
	tasks = append(tasks, Task{Title: title, Complete: false, Description: Description, Priority: Priority})
	saveTasksToJsonFile(Filename)
}

func completeTask(index int) {

	if index >= 0 && index < len(tasks) {
		tasks[index].Complete = true
		saveTasksToJsonFile(Filename)
	}
}

func deleteTask(index int) {
	if index >= 0 && index < len(tasks) {
		fmt.Println(tasks[:index], tasks[index+1:])
		tasks = append(tasks[:index], tasks[index+1:]...)
		saveTasksToJsonFile(Filename)
	}
}

func showTasks() {
	//对tasks进行排序， 使用Priority字段进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})
	for i, task := range tasks {
		status := "未完成"
		if task.Complete {
			status = "已完成"
		}
		fmt.Printf("%d. [%s] %s - %s 优先级: %d \n", i+1, status, task.Title, task.Description, task.Priority)

	}
}
//显示已完成的任务
func showCompletedTasks() {
	for i, task := range tasks {
		if task.Complete {
			fmt.Printf("%d. %s \n", i+1, task.Title)
		}
	}
	// filter Completed tasks to new slice

}

func (t Task) showTaskComplete() {
	var status string
	if t.Complete {
		status = "已完成"
	} else {
		status = "未完成"
	}
	fmt.Printf("%s【%s】%s \n", "任务： ", t.Title, status)
}

func saveTasksToJsonFile(filename string) {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(filename, data, 0666)
}
