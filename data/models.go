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

func (t TaskType) Int() int64 {
	return int64(t)
}

func TaskTypeFrom(t int64) TaskType {
	switch t {
	case int64(InboxTask):
		return InboxTask
	case int64(ActionTask):
		return ActionTask
	case int64(SomedayTask):
		return SomedayTask
	case int64(TicklerTask):
		return TicklerTask
	default:
		return InboxTask
	}
}

type Task struct {
	Id      Id
	Name    string
	Done    time.Time
	Created time.Time
	Type    TaskType
}

func ParseUnixDate(sec interface{}) time.Time {
	if t, ok := sec.(int64); ok {
		return time.Unix(t, 0)
	} else {
		return time.Time{}
	}
}
