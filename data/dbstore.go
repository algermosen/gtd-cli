package data

type dbStore struct{}

func NewDbStore() Store {
	return &dbStore{}
}

func (s *dbStore) Inbox() []Task {
	return nil
}

func (s *dbStore) Actions() []Task {
	return nil
}

func (s *dbStore) Somedays() []Task {
	return nil
}

func (s *dbStore) Ticklers() []Task {
	return nil
}

func (s *dbStore) Add(Task) error {
	return nil
}

func (s *dbStore) Delete(Id) error {
	return nil
}

func (s *dbStore) Check(Id) error {
	return nil
}

func (s *dbStore) CheckAll() error {
	return nil
}

func (s *dbStore) DeleteAll() error {
	return nil
}
