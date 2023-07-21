package database

type DataBase interface {
	Set(json string) error
	List() ([]string, error)
	Close() error
}
