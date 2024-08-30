package data

type fileStore struct{}

func NewFileStore() Store {
	return &fileStore{}
}

func (s *fileStore) Inbox() []Task {
	return nil
}

func (s *fileStore) Actions() []Task {
	return nil
}

func (s *fileStore) Somedays() []Task {
	return nil
}

func (s *fileStore) Ticklers() []Task {
	return nil
}

func (s *fileStore) Add(Task) error {
	return nil
}

func (s *fileStore) Delete(Id) error {
	return nil
}

func (s *fileStore) Check(Id) error {
	return nil
}

func (s *fileStore) CheckAll() error {
	return nil
}

func (s *fileStore) DeleteAll() error {
	return nil
}
