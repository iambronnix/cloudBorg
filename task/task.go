package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type state int

const(
	Pending state = iota
	Scheduled
	Running 
	Completed
	Failed
)

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

type TaskEvent struct{
	ID uuid.UUID
	State state
	TimeStamp time.Time
	Task Task
}