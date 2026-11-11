package worker

import (
	"cloudBorg/task"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

//worker requirements:
// Run tasks as Docker containers
// Accept tasks to run from a manager
// Provide relevant statistics to the manager for the purpose of scheduling tasks
// keep track of its tasks and their state

type Worker struct{
	Name string 
	Queue queue.Queue
	Db map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStat(){
	fmt.Println("i'll collect task stats")
}

func (w *Worker) RunTask(){
	fmt.Println("I'll start or stop a task")
}

func (w *Worker) StartTask(){
	fmt.Println("i'll start a task")
}

func (w *Worker) StopTask(){
	fmt.Println("i'll stop a task")
}