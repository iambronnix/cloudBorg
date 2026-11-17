package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)
//defines states of a given task
type state int

const(
	Pending state = iota
	Scheduled
	Running 
	Completed
	Failed
)

//attributes of a task
type Task struct{
	ID uuid.UUID
	Name string
	State state
	Image string
	Memory int
	Disk int
	ExposedPorts nat.PortSet
	PortBindings map[string]string
	RestartPolicy string
	StartTime time.Time
	FinishTime time.Time
}

//used by a user to track a task
type TaskEvent struct{
	ID uuid.UUID
	State state
	TimeStamp time.Time
	Task Task
}