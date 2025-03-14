package more

type IRepository interface {
	Insert(string) bool
	Delete(string) bool
}

type MysqlRepository struct {
	Items map[string]bool
}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{
		Items: make(map[string]bool, 0),
	}
}

type PostgresRepository struct {
	Items map[string]bool
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{
		Items: make(map[string]bool, 0),
	}
}

func (repo *MysqlRepository) Insert(name string) bool {
	repo.Items[name] = true
	return true
}
func (repo *MysqlRepository) Delete(name string) bool {
	delete(repo.Items, name)
	return true
}

func (repo *PostgresRepository) Insert(name string) bool {
	repo.Items[name] = true
	return true
}
func (repo *PostgresRepository) Delete(name string) bool {
	delete(repo.Items, name)
	return true
}
