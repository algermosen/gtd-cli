package data

type sqliteStore struct{}

func (s *sqliteStore) Inbox() []Task {
	return nil
}

func (s *sqliteStore) Actions() []Task {
	return nil
}

func (s *sqliteStore) Somedays() []Task {
	return nil
}

func (s *sqliteStore) Ticklers() []Task {
	return nil
}

func (s *sqliteStore) Add(Task) error {
	return nil
}

func (s *sqliteStore) Delete(Id) error {
	return nil
}

func (s *sqliteStore) Check(Id) error {
	return nil
}

func (s *sqliteStore) CheckAll() error {
	return nil
}

func (s *sqliteStore) DeleteAll() error {
	return nil
}
