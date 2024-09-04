package data

import (
	"context"
	"database/sql"
	database "gtd/data/sqlite"
	"log"
)

func NewSqliteStore(path string) Store {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	q := database.New(db)
	return &sqliteStore{
		Queries: q,
	}
}

type sqliteStore struct {
	*database.Queries
}

var emptyTasks = make([]Task, 0)

func (s *sqliteStore) Inbox() []Task {
	tasks, err := s.ListTasks(context.Background(), int64(InboxTask))

	if err != nil {
		log.Print(err)
		return emptyTasks
	}
	return parseTasks(tasks)
}

func (s *sqliteStore) Actions() []Task {
	tasks, err := s.ListTasks(context.Background(), int64(ActionTask))

	if err != nil {
		log.Print(err)
		return emptyTasks
	}
	return parseTasks(tasks)
}

func (s *sqliteStore) Somedays() []Task {
	tasks, err := s.ListTasks(context.Background(), int64(SomedayTask))

	if err != nil {
		log.Print(err)
		return emptyTasks
	}
	return parseTasks(tasks)
}

func (s *sqliteStore) Ticklers() []Task {
	tasks, err := s.ListTasks(context.Background(), int64(TicklerTask))

	if err != nil {
		log.Print(err)
		return emptyTasks
	}
	return parseTasks(tasks)
}

func (s *sqliteStore) Add(task *Task) error {
	params := database.CreateTaskParams{
		Name:    task.Name,
		Done:    task.Done.Unix(),
		Created: task.Created.Unix(),
		Type:    task.Type.Int(),
	}
	id, err := s.CreateTask(context.Background(), params)

	if err != nil {
		return err
	}

	task.Id = Id(id)
	return nil
}

func (s *sqliteStore) Delete(id Id) error {
	err := s.DeleteTask(context.Background(), int64(id))
	return err
}

func (s *sqliteStore) Check(id Id) error {
	err := s.CheckTask(context.Background(), int64(id))
	return err
}

func (s *sqliteStore) CheckAll() error {
	err := s.CheckAllTasks(context.Background())
	return err
}

func (s *sqliteStore) DeleteAll() error {
	err := s.DeleteAllTasks(context.Background())
	return err
}

func parseTasks(tasks []database.Task) []Task {
	var parsed = make([]Task, len(tasks))
	for i, task := range tasks {
		parsed[i] = Task{
			Id:      Id(task.ID),
			Name:    task.Name,
			Done:    ParseUnixDate(task.Done),
			Created: ParseUnixDate(task.Created),
			Type:    TaskTypeFrom(task.Type),
		}
	}
	return parsed
}
