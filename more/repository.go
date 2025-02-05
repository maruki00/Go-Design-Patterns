package main

type IRepository interface {
	Insert(string) bool
	Delete(string) bool
}

type MysqlRepository struct {
	items map[string]bool
}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{
		items: make(map[string]bool, 0),
	}
}

type PostgresRepository struct {
	items map[string]bool
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{
		items: make(map[string]bool, 0),
	}
}

func (repo *MysqlRepository) Insert(name string) bool {
	repo.items[name] = true
	return true
}
func (repo *MysqlRepository) Delete(name string) bool {
	delete(repo.items, name)
	return true
}

func (repo *PostgresRepository) Insert(name string) bool {
	repo.items[name] = true
	return true
}

func (repo *PostgresRepository) Delete(name string) bool {
	delete(repo.items, name)
	return true
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(name string) bool {
	return s.repo.Insert(name)
}

func main() {
	service := NewService(NewMysqlRepository())
	service.Create("hello world")
}
