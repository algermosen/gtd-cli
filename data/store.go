package data

import (
	"time"
)

type TaskType int
type Id uint16

const (
	InboxTask TaskType = iota
	ActionTask
	SomedayTask
	TicklerTask
)

type Task struct {
	Id      Id
	Name    string
	Done    time.Time
	Created time.Time
	Type    TaskType
}

type Store interface {
	Inbox() []Task
	Actions() []Task
	Somedays() []Task
	Ticklers() []Task
	Add(Task) error
	Delete(Id) error
	Check(Id) error
	CheckAll() error
	DeleteAll() error
}
