package cache

type Cache interface {
	Set(json string) error
	SetList(jsonList []string) error
	Single(kay string) (string, error)
}
