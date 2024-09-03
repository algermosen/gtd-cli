package data

import "time"

type TaskType int
type Id uint16

const (
	InboxTask TaskType = iota
	ActionTask
	SomedayTask
	TicklerTask
	DefaultTask = InboxTask
)

type Task struct {
	Id      Id
	Name    string
	Done    time.Time
	Created time.Time
	Type    TaskType
}
