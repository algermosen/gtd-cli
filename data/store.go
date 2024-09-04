package data

type Store interface {
	Inbox() []Task
	Actions() []Task
	Somedays() []Task
	Ticklers() []Task
	Add(string) error
	Delete(Id) error
	Check(Id) error
	CheckAll() error
	DeleteAll() error
}
