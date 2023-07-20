package cache

type Cache interface {
	Set(Payload) error
	SetList([]Payload) error
	Single(kay string) (Payload, error)
}

type Payload struct {
	key   string
	value []byte
}

func NewPayload(key string, value []byte) Payload {
	return Payload{
		key:   key,
		value: value,
	}
}

func (p Payload) Key() string {
	return p.key
}

func (p Payload) Value() []byte {
	return p.value
}
